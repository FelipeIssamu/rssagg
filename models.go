package main

import (
	"time"

	"github.com/FelipeIssamu/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
    ID uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Name string `json:"name"`
    APIKey string `json:"api_key"`
}

type Feed struct {
    ID uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Name string `json:"name"`
    Url string `json:"url"`
    UserID uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
    ID uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    UserID uuid.UUID `json:"user_id"`
    FeedID uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
    return User{
        ID: dbUser.ID,
        CreatedAt: dbUser.CreatedAt,
        UpdatedAt: dbUser.UpdatedAt,
        Name: dbUser.Name,
        APIKey: dbUser.ApiKey,
    }
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
    return Feed{
        ID: dbFeed.ID,
        CreatedAt: dbFeed.CreatedAt,
        UpdatedAt: dbFeed.UpdatedAt,
        Name: dbFeed.Name,
        Url: dbFeed.Url,
        UserID: dbFeed.UserID,
    }
}

func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
    feeds := []Feed{}
    for _, dbFeed := range dbFeed {
        feeds = append(feeds, databaseFeedToFeed(dbFeed))
    }
    return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
    return FeedFollow{
        ID: dbFeedFollow.ID,
        CreatedAt: dbFeedFollow.CreatedAt,
        UpdatedAt: dbFeedFollow.UpdatedAt,
        UserID: dbFeedFollow.UserID,
        FeedID: dbFeedFollow.FeedID,
    }
}
