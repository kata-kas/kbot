package kajabi

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/klubbot/klubbot/internal/normalization"
)

type User struct {
	Name string
}

func ScrapeUsers(url string) []string {
	users := getAllSubscribedUsers(url)

	formattedUsers := make([]string, len(users))

	for idx, user := range users {
		formattedName := normalization.NormalizeForComparison(user.Name)
		formattedUsers[idx] = formattedName
	}

	return formattedUsers
}

func IsUserSubbed(name string) bool {
	path, _ := launcher.LookPath()
	u := launcher.New().Set("--no-sandbox", "true").Bin(path).NoSandbox(true).MustLaunch()
	browser := rod.New().ControlURL(u).Timeout(time.Hour).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser)
	login(page)
	page.MustNavigate(os.Getenv("SEARCH_URL") + name)
	page.MustWaitLoad()
	page.MustWaitDOMStable()
	page.MustScreenshot("search.png")
	emptyPageClass := ".sage-empty-state__title"
	if page.MustHas(emptyPageClass) {
		fmt.Printf("User %s not found\n", name)
		return false
	}
	fmt.Printf("User %s found\n", name)
	return true
}

func login(page *rod.Page) {
	page.MustNavigate(os.Getenv("LOGIN_URL"))
	page.MustWaitDOMStable()

	if !page.MustHas("form") {
		return
	}

	page.MustElement("form")
	page.MustScreenshot("")
	page.MustElement("#username").MustClick().MustInput(os.Getenv("USERNAME"))
	page.MustScreenshot("")
	page.MustElement("#password").MustClick().MustInput(os.Getenv("PASSWORD"))
	page.MustScreenshot("")
	page.MustElement(".kjb-submit-btn").MustClick()
	page.MustScreenshot("")
	page.MustWaitLoad()
	page.MustScreenshot("")
}

func getAllSubscribedUsers(url string) []User {
	path, _ := launcher.LookPath()
	u := launcher.New().Set("--no-sandbox", "true").Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).Timeout(time.Hour).Trace(true).MustConnect()
	defer browser.MustClose()

	pool := rod.NewPagePool(5)

	create := func() *rod.Page {
		return stealth.MustPage(browser).Browser().MustIncognito().MustPage()
	}
	page := pool.Get(create)
	totalNumberOfPages := getTotalNumberOfPages(page, url)
	pool.Put(page)

	totalUsers := make([]User, 0)
	extractUsers := func(page *rod.Page, i int) {
		login(page)
		newURL := fmt.Sprintf("%s&page=%s", url, strconv.Itoa(i))
		fmt.Println(newURL)
		page.MustNavigate(newURL)
		page.Race().Element("table").MustHandle(func(e *rod.Element) {
			page.MustScreenshot(fmt.Sprintf("page-%d.png", i))
			elements := e.MustElements(".contacts-table-details-name__text")
			fmt.Printf("Number of users: %d\n", len(elements))
			users := make([]User, len(elements))
			for i, el := range elements {
				users[i] = User{Name: el.MustText()}
			}
			totalUsers = append(totalUsers, users...)
			fmt.Printf("Total users: %d\n", len(totalUsers))
		}).MustDo()
	}

	var wg sync.WaitGroup
	for i := 1; i <= totalNumberOfPages; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			page := pool.Get(create)
			defer pool.Put(page)
			extractUsers(page, i)
			fmt.Printf("Finished page %d\n", i)
		}(i)
	}
	wg.Wait()

	pool.Cleanup(func(p *rod.Page) {
		p.MustClose()
	})

	fmt.Println("Finished scraping users")
	return totalUsers
}

func getTotalNumberOfPages(page *rod.Page, url string) int {
	login(page)
	page.MustNavigate(url)
	page.MustElement("table")
	paginationItems := page.MustElements(".sage-pagination__item")
	numberOfPagesStr := paginationItems[len(paginationItems)-2].MustText()
	fmt.Printf("Number of pages: %s\n", numberOfPagesStr)
	numberOfPages, err := strconv.ParseInt(numberOfPagesStr, 10, 10)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return int(numberOfPages)
}
