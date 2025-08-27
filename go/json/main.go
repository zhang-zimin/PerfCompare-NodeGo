package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Simple data structure
type SimpleData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

// Medium complexity data structure
type Profile struct {
	Age       int      `json:"age"`
	City      string   `json:"city"`
	Interests []string `json:"interests"`
	Settings  struct {
		Theme         string `json:"theme"`
		Notifications bool   `json:"notifications"`
		Privacy       struct {
			Public  bool `json:"public"`
			Friends bool `json:"friends"`
		} `json:"privacy"`
	} `json:"settings"`
}

type OrderItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	ID      int         `json:"id"`
	Product string      `json:"product"`
	Price   float64     `json:"price"`
	Date    string      `json:"date"`
	Items   []OrderItem `json:"items"`
}

type MediumData struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Profile Profile `json:"profile"`
	Orders  []Order `json:"orders"`
}

// Complex data structure
type SocialData struct {
	Twitter  string `json:"twitter"`
	GitHub   string `json:"github"`
	LinkedIn string `json:"linkedin"`
}

type Notifications struct {
	Email bool `json:"email"`
	Push  bool `json:"push"`
	SMS   bool `json:"sms"`
}

type Preferences struct {
	Theme         string        `json:"theme"`
	Language      string        `json:"language"`
	Notifications Notifications `json:"notifications"`
}

type ActivityMetadata struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	SessionID string `json:"session_id"`
}

type Activity struct {
	ID        int              `json:"id"`
	Type      string           `json:"type"`
	Timestamp string           `json:"timestamp"`
	Metadata  ActivityMetadata `json:"metadata"`
}

type User struct {
	ID          int         `json:"id"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Profile     UserProfile `json:"profile"`
	Preferences Preferences `json:"preferences"`
	Activity    []Activity  `json:"activity"`
}

type UserProfile struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Bio       string     `json:"bio"`
	Avatar    string     `json:"avatar"`
	Social    SocialData `json:"social"`
}

type ComplexData struct {
	Metadata struct {
		Version   string `json:"version"`
		Generated string `json:"generated"`
		Schema    string `json:"schema"`
	} `json:"metadata"`
	Users []User `json:"users"`
}

func generateSimpleData() SimpleData {
	return SimpleData{
		ID:     1,
		Name:   "Test User",
		Email:  "test@example.com",
		Active: true,
	}
}

func generateMediumData() MediumData {
	orders := make([]Order, 100)
	for i := 0; i < 100; i++ {
		items := make([]OrderItem, (i%5)+1)
		for j := 0; j < len(items); j++ {
			items[j] = OrderItem{
				ID:       j,
				Name:     fmt.Sprintf("Item %d", j),
				Quantity: (j % 10) + 1,
			}
		}
		orders[i] = Order{
			ID:      i,
			Product: fmt.Sprintf("Product %d", i),
			Price:   float64(i) * 1.5,
			Date:    time.Now().Format(time.RFC3339),
			Items:   items,
		}
	}

	return MediumData{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
		Profile: Profile{
			Age:       30,
			City:      "New York",
			Interests: []string{"coding", "music", "travel"},
			Settings: struct {
				Theme         string `json:"theme"`
				Notifications bool   `json:"notifications"`
				Privacy       struct {
					Public  bool `json:"public"`
					Friends bool `json:"friends"`
				} `json:"privacy"`
			}{
				Theme:         "dark",
				Notifications: true,
				Privacy: struct {
					Public  bool `json:"public"`
					Friends bool `json:"friends"`
				}{
					Public:  false,
					Friends: true,
				},
			},
		},
		Orders: orders,
	}
}

func generateComplexData() ComplexData {
	users := make([]User, 1000)
	for i := 0; i < 1000; i++ {
		activities := make([]Activity, 50)
		for j := 0; j < 50; j++ {
			activities[j] = Activity{
				ID:        j,
				Type:      []string{"login", "logout", "purchase", "view"}[j%4],
				Timestamp: time.Now().Add(-time.Duration(j) * time.Hour).Format(time.RFC3339),
				Metadata: ActivityMetadata{
					IP:        fmt.Sprintf("192.168.1.%d", j%255),
					UserAgent: fmt.Sprintf("Browser %d", j%10),
					SessionID: fmt.Sprintf("session-%d-%d", i, j),
				},
			}
		}

		users[i] = User{
			ID:       i,
			Username: fmt.Sprintf("user%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Profile: UserProfile{
				FirstName: fmt.Sprintf("First%d", i),
				LastName:  fmt.Sprintf("Last%d", i),
				Bio:       fmt.Sprintf("This is a bio for user %d. ", i) + strings.Repeat("Bio content. ", 10),
				Avatar:    fmt.Sprintf("https://example.com/avatar/%d.jpg", i),
				Social: SocialData{
					Twitter:  fmt.Sprintf("@user%d", i),
					GitHub:   fmt.Sprintf("user%d", i),
					LinkedIn: fmt.Sprintf("user-%d", i),
				},
			},
			Preferences: Preferences{
				Theme:    []string{"light", "dark"}[i%2],
				Language: []string{"en", "zh", "es", "fr"}[i%4],
				Notifications: Notifications{
					Email: i%2 == 0,
					Push:  i%3 == 0,
					SMS:   i%5 == 0,
				},
			},
			Activity: activities,
		}
	}

	return ComplexData{
		Metadata: struct {
			Version   string `json:"version"`
			Generated string `json:"generated"`
			Schema    string `json:"schema"`
		}{
			Version:   "1.0.0",
			Generated: time.Now().Format(time.RFC3339),
			Schema:    "user-data-v1",
		},
		Users: users,
	}
}

func testSerialization(data interface{}, iterations int) (float64, string) {
	fmt.Printf("Testing JSON serialization...\n")
	var times []float64
	var jsonString string

	for i := 0; i < iterations; i++ {
		start := time.Now()
		jsonBytes, err := json.Marshal(data)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("Serialization error: %v\n", err)
			continue
		}

		jsonString = string(jsonBytes)
		timeMs := float64(elapsed.Nanoseconds()) / 1000000.0
		times = append(times, timeMs)
		fmt.Printf("Serialization %d: %.3fms\n", i+1, timeMs)
	}

	var total float64
	for _, t := range times {
		total += t
	}
	avgTime := total / float64(len(times))
	fmt.Printf("Average serialization time: %.3fms\n", avgTime)
	fmt.Printf("JSON size: %d characters\n\n", len(jsonString))

	return avgTime, jsonString
}

func testDeserialization(jsonString string, target interface{}, iterations int) float64 {
	fmt.Printf("Testing JSON deserialization...\n")
	var times []float64

	for i := 0; i < iterations; i++ {
		start := time.Now()
		err := json.Unmarshal([]byte(jsonString), target)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("Deserialization error: %v\n", err)
			continue
		}

		timeMs := float64(elapsed.Nanoseconds()) / 1000000.0
		times = append(times, timeMs)
		fmt.Printf("Deserialization %d: %.3fms\n", i+1, timeMs)
	}

	var total float64
	for _, t := range times {
		total += t
	}
	avgTime := total / float64(len(times))
	fmt.Printf("Average deserialization time: %.3fms\n\n", avgTime)

	return avgTime
}

func testRoundTrip(data interface{}, target interface{}, iterations int) float64 {
	fmt.Printf("Testing round-trip (serialize + deserialize)...\n")
	var times []float64

	for i := 0; i < iterations; i++ {
		start := time.Now()
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Serialization error: %v\n", err)
			continue
		}

		err = json.Unmarshal(jsonBytes, target)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("Deserialization error: %v\n", err)
			continue
		}

		timeMs := float64(elapsed.Nanoseconds()) / 1000000.0
		times = append(times, timeMs)
		fmt.Printf("Round-trip %d: %.3fms\n", i+1, timeMs)
	}

	var total float64
	for _, t := range times {
		total += t
	}
	avgTime := total / float64(len(times))
	fmt.Printf("Average round-trip time: %.3fms\n\n", avgTime)

	return avgTime
}

func main() {
	fmt.Println("Go JSON Processing Tests")
	fmt.Println("========================")

	// Simple data test
	fmt.Printf("\n=== SIMPLE Data Structure ===\n")
	simpleData := generateSimpleData()
	serTime, jsonStr := testSerialization(simpleData, 5)
	var simpleTarget SimpleData
	deserTime := testDeserialization(jsonStr, &simpleTarget, 5)
	var simpleRoundTripTarget SimpleData
	roundTripTime := testRoundTrip(simpleData, &simpleRoundTripTarget, 5)

	fmt.Printf("Summary for simple data:\n")
	fmt.Printf("- Serialization: %.3fms\n", serTime)
	fmt.Printf("- Deserialization: %.3fms\n", deserTime)
	fmt.Printf("- Round-trip: %.3fms\n", roundTripTime)
	fmt.Println("---")

	// Medium data test
	fmt.Printf("\n=== MEDIUM Data Structure ===\n")
	mediumData := generateMediumData()
	serTime, jsonStr = testSerialization(mediumData, 5)
	var mediumTarget MediumData
	deserTime = testDeserialization(jsonStr, &mediumTarget, 5)
	var mediumRoundTripTarget MediumData
	roundTripTime = testRoundTrip(mediumData, &mediumRoundTripTarget, 5)

	fmt.Printf("Summary for medium data:\n")
	fmt.Printf("- Serialization: %.3fms\n", serTime)
	fmt.Printf("- Deserialization: %.3fms\n", deserTime)
	fmt.Printf("- Round-trip: %.3fms\n", roundTripTime)
	fmt.Println("---")

	// Complex data test
	fmt.Printf("\n=== COMPLEX Data Structure ===\n")
	complexData := generateComplexData()
	serTime, jsonStr = testSerialization(complexData, 5)
	var complexTarget ComplexData
	deserTime = testDeserialization(jsonStr, &complexTarget, 5)
	var complexRoundTripTarget ComplexData
	roundTripTime = testRoundTrip(complexData, &complexRoundTripTarget, 5)

	fmt.Printf("Summary for complex data:\n")
	fmt.Printf("- Serialization: %.3fms\n", serTime)
	fmt.Printf("- Deserialization: %.3fms\n", deserTime)
	fmt.Printf("- Round-trip: %.3fms\n", roundTripTime)
	fmt.Println("---")
}
