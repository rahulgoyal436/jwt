
// ********RoostGPT********
/*

roost_feedback [04/07/2025, 1:20:57 PM]:enhance and improve the test\n\n

roost_feedback [08/07/2025, 12:25:39 PM]:enhance the test feature file\n\n

roost_feedback [08/07/2025, 12:41:22 PM]:enhance the feature file \n\n
*/

// ********RoostGPT********

Feature: YouTube API Testing

Background:
  Given the API base URL is "https://www.googleapis.com/youtube/v3"
  And the authorization header is set with a valid API key
  And the content type is "application/json"
  And the request timeout is set to 30 seconds
  And the response format is set to JSON
  And the request logging is enabled for debugging purposes
  And error responses are captured for analysis

Scenario: Search for videos by keyword
  When I send a GET request to "/search" with the following parameters:
    | part             | snippet       |
    | q                | test automation |
    | maxResults       | 5             |
    | type             | video         |
    | relevanceLanguage | en           |
    | safeSearch       | moderate      |
  Then the response status code should be 200
  And the response should contain a list of videos
  And each video in the response should have an ID and title
  And each video should have a valid thumbnail URL
  And the response time should be less than 2 seconds
  And the response should contain exactly 5 items
  And the response headers should contain valid cache control directives
  And the response should be properly formatted according to the YouTube API schema

Scenario: Get video details by ID
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet,statistics,contentDetails,player,status |
    | id         | dQw4w9WgXcQ                                   |
  Then the response status code should be 200
  And the response should contain video details
  And the video details should include title, description, and view count
  And the video duration should be in ISO 8601 format
  And the response should include the channel information
  And the response should include the video's category ID
  And the video statistics should include like count and comment count
  And the video should have a valid embed HTML
  And the video's privacy status should be "public"

Scenario: Handle invalid API key
  Given the authorization header is set with an invalid API key "AIzaSyInvalidKeyExample123456789"
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
  Then the response status code should be 403
  And the response should contain an error message about invalid credentials
  And the error reason should be "API key not valid"
  And the error message should contain "API key not valid. Please pass a valid API key"
  And the response should include a reference to the API documentation
  And the response time should be less than 1 second

Scenario: Verify pagination of search results
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 10            |
    | type       | video         |
    | order      | relevance     |
  Then the response status code should be 200
  And the response should contain a nextPageToken
  And I store the video IDs from the response
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 10            |
    | type       | video         |
    | pageToken  | {nextPageToken} |
    | order      | relevance     |
  Then the response status code should be 200
  And the response should contain different videos than the first page
  And the total results count should be greater than the maxResults parameter
  And the response should contain a valid prevPageToken matching the first request
  And navigating through 3 pages should return unique results

Scenario: Filter videos by publish date
  When I send a GET request to "/search" with the following parameters:
    | part           | snippet       |
    | q              | test automation |
    | maxResults     | 5             |
    | type           | video         |
    | publishedAfter | 2022-01-01T00:00:00Z |
    | publishedBefore| 2023-01-01T00:00:00Z |
  Then the response status code should be 200
  And all videos in the response should have been published after "2022-01-01"
  And all videos in the response should have been published before "2023-01-01"
  And the response should contain a list of videos
  And the videos should be sorted by relevance by default

Scenario: Get channel information
  When I send a GET request to "/channels" with the following parameters:
    | part       | snippet,statistics,contentDetails,brandingSettings |
    | id         | UC_x5XG1OV2P6uZZ5FSM9Ttw                         |
  Then the response status code should be 200
  And the response should contain channel details
  And the channel details should include title, description, and subscriber count
  And the response should include the channel's playlist IDs
  And the channel statistics should include view count and video count
  And the channel should have valid thumbnail images
  And the channel's branding settings should include a banner image URL
  And the channel's country of origin should be specified

Scenario: Handle resource not found
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet       |
    | id         | nonExistentVideoId123 |
  Then the response status code should be 200
  And the response should contain an empty items array
  And the response should include a valid kind property
  And the response should include a valid etag
  And the response should not contain any error messages

Scenario: Rate limiting behavior
  When I send 10 consecutive GET requests to "/search" with minimal delay
  Then all responses should have status code 200 or 429
  And if status code is 429, the response should contain quota exceeded information
  And the response headers should include rate limiting information
  And the quota usage information should be present in the response
  And the error response should include a retry-after header if rate limited
  And the error message should provide guidance on quota management

Scenario: Search videos by category
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 5             |
    | type       | video         |
    | videoCategoryId | 28       |
  Then the response status code should be 200
  And all videos in the response should belong to category ID 28
  And the response should contain a list of videos
  And each video should have relevant tags related to the search query

Scenario: Get related videos
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | relatedToVideoId | dQw4w9WgXcQ |
    | maxResults | 5             |
    | type       | video         |
  Then the response status code should be 200
  And the response should contain a list of related videos
  And the videos should be topically related to the specified video
  And none of the videos should have the same ID as the relatedToVideoId

Scenario: Search for live streams
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | eventType  | live          |
    | maxResults | 5             |
    | type       | video         |
  Then the response status code should be 200
  And all videos in the response should have a live broadcast content status
  And the response should include live streaming details for each video