package db

import (
	"SocialMedia/internal/models"
	"SocialMedia/internal/store"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
)

var usernames = []string{
	"Tonny", "RockyBoy", "NinjaStar", "AceHigh", "PiratePete", "SpaceCadet", "CodeMaster", "GameOn", "Sparky22", "ByteBandit", "DigitalDude", "CyberSquad", "GeekyGuru", "TechTitan", "PixelPerfect", "VirusHunter", "MalwareMVP", "CodeCrafter", "WebWizard", "APIAce", "DatabaseDrill", "DevilishDave", "SuperStar99", "RocketMan2000", "MoonlightMike", "SolarFlare3000", "GalacticGamer", "AstroNerd123", "CosmicCoder", "QuantumQuester", "NebulaNemesis", "AuroraAvatar", "LunaLovegood", "StellarSavant", "CelestialCoder", "OrionOutlaw", "PulsarPwnz", "BlazingByte", "FuriousFlux", "RapidRouter", "MaverickMage", "VaporVelocity", "StormySwift", "ThunderousTitan", "LightningLapse", "SonicSpeedster", "TurboTornado", "HurricaneHacker", "CycloneCoder", "TyphoonTechnician",
}

var titles = []string{
	"Tech Tips", "Coding Secrets", "Dev Diary", "Geeky Gadgets", "Code Confessions", "Web Wizardry", "Digital Dreams", "Gaming Guide", "Byte-Sized Wisdom", "Pixel Perfect", "Algorithmic Adventures", "Cyber Chronicles", "The Code Cave", "Tech Tales", "Software Scoop", "Programming Ponderings", "Web Wonders", "Code Crafted", "Digital Diary", "Gadget Guru",
}

var contents = []string{
	"Tech Tips: Learn how to optimize your code for better performance and efficiency.",
	"Coding Secrets: Discover hidden gems and tricks of the trade to take your coding skills to the next level.",
	"Dev Diary: Join me on my journey as a developer, where I share my experiences, successes, and failures.",
	"Geeky Gadgets: Get the latest news and reviews on the coolest gadgets and technology trends.",
	"Code Confessions: Share your coding struggles and triumphs, and get advice from fellow developers.",
	"Web Wizardry: Learn how to create stunning web designs and user experiences with our expert tutorials.",
	"Digital Dreams: Explore the latest developments in AI, machine learning, and data science.",
	"Gaming Guide: Get tips, tricks, and strategies for your favorite games, as well as reviews of new releases.",
	"Byte-Sized Wisdom: Share bite-sized pieces of coding wisdom, humor, and inspiration.",
	"Pixel Perfect: Learn how to create stunning visuals with our expert tutorials on graphics design and animation.",
	"Algorithmic Adventures: Embark on a journey through the world of algorithms, data structures, and software engineering.",
	"Cyber Chronicles: Stay up-to-date with the latest news and trends in cybersecurity, online safety, and digital rights.",
	"The Code Cave: A community-driven blog where developers share their experiences, knowledge, and ideas.",
	"Tech Tales: Read inspiring stories of innovation, entrepreneurship, and technological advancements.",
	"Software Scoop: Get the latest scoop on software development trends, technologies, and best practices.",
	"Programming Ponderings: Explore the intersection of programming, philosophy, and human experience.",
	"Web Wonders: Marvel at the magic of web design, user experience, and digital storytelling.",
	"Code Crafted: Learn how to craft beautiful code, from architecture to implementation, with our expert tutorials.",
	"Digital Diary: Join me on my journey as a developer, where I share my experiences, successes, and failures.",
	"Gadget Guru: Get the latest news and reviews on the coolest gadgets and technology trends.",
}

var tags = []string{
	"tech", "coding", "development", "programming", "software", "engineering",
	"gaming", "game-dev", "graphics", "animation", "ai", "machine-learning",
	"data-science", "cybersecurity", "online-safety", "digital-rights",
	"web-design", "user-experience", "ux", "ui", "design",
	"innovation", "entrepreneurship", "tech-trends", "programming-languages",
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I completely agree with your thoughts.",
	"Thanks for the tips, very helpful.",
	"Interesting perspective, I hadn't considered that.",
	"Thanks for sharing your experience.",
	"Well written, I enjoyed reading this.",
	"This is very insightful, thanks for posting.",
	"Great advice, I'll definitely try that.",
	"I love this, very inspirational.",
	"Thanks for the information, very useful.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user: ", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating posts: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comments: ", err)
			return
		}
	}

	log.Println("Database seeding completed.")
}

func generateUsers(num int) []*models.User {
	users := make([]*models.User, num)

	for i := 0; i < num; i++ {
		users[i] = &models.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Role: models.Role{
				Name: "user",
			},
		}
		if err := users[i].Password.Set("12345"); err != nil {
			log.Fatal("Error hashing a password")
		}
	}

	return users
}

func generatePosts(num int, users []*models.User) []*models.Post {
	posts := make([]*models.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &models.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*models.User, posts []*models.Post) []*models.Comment {
	cms := make([]*models.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &models.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
