
// ********RoostGPT********
/*

roost_feedback [04/07/2025, 1:20:57 PM]:enhance and improve the test\n\n

roost_feedback [08/07/2025, 12:25:39 PM]:enhance the test feature file\n\n
*/

// ********RoostGPT********

Feature: YouTube API Testing

Background:
  Given the API base URL is "https://www.googleapis.com/youtube/v3"
  And the authorization header is set with a valid API key
  And the content type is "application/json"
  And the request timeout is set to 30 seconds
  And the response format is set to JSON

Scenario: Search for videos by keyword
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 5             |
    | type       | video         |
    | relevanceLanguage | en     |
  Then the response status code should be 200
  And the response should contain a list of videos
  And each video in the response should have an ID and title
  And the response time should be less than 2 seconds
  And the response should contain exactly 5 items

Scenario: Get video details by ID
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet,statistics,contentDetails |
    | id         | dQw4w9WgXcQ                      |
  Then the response status code should be 200
  And the response should contain video details
  And the video details should include title, description, and view count
  And the video duration should be in ISO 8601 format
  And the response should include the channel information
  And the response should include the video's category ID

Scenario: Handle invalid API key
  Given the authorization header is set with an invalid API key "AIzaSyInvalidKeyExample123456789"
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
  Then the response status code should be 403
  And the response should contain an error message about invalid credentials
  And the error reason should be "API key not valid"
  And the error message should contain "API key not valid. Please pass a valid API key"

Scenario: Verify pagination of search results
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 10            |
    | type       | video         |
  Then the response status code should be 200
  And the response should contain a nextPageToken
  And I store the video IDs from the response
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 10            |
    | type       | video         |
    | pageToken  | {nextPageToken} |
  Then the response status code should be 200
  And the response should contain different videos than the first page
  And the total results count should be greater than the maxResults parameter

Scenario: Filter videos by publish date
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 5             |
    | type       | video         |
    | publishedAfter | 2022-01-01T00:00:00Z |
  Then the response status code should be 200
  And all videos in the response should have been published after "2022-01-01"
  And the response should contain a list of videos

Scenario: Get channel information
  When I send a GET request to "/channels" with the following parameters:
    | part       | snippet,statistics,contentDetails |
    | id         | UC_x5XG1OV2P6uZZ5FSM9Ttw         |
  Then the response status code should be 200
  And the response should contain channel details
  And the channel details should include title, description, and subscriber count
  And the response should include the channel's playlist IDs

Scenario: Handle resource not found
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet       |
    | id         | nonExistentVideoId123 |
  Then the response status code should be 200
  And the response should contain an empty items array

Scenario: Rate limiting behavior
  When I send 10 consecutive GET requests to "/search" with minimal delay
  Then all responses should have status code 200 or 429
  And if status code is 429, the response should contain quota exceeded information
  And the response headers should include rate limiting information