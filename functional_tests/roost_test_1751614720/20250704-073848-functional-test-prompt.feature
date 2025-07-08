
// ********RoostGPT********
/*

roost_feedback [04/07/2025, 1:20:57 PM]:enhance and improve the test\n\n

roost_feedback [08/07/2025, 12:25:39 PM]:enhance the test feature file\n\n

roost_feedback [08/07/2025, 12:41:22 PM]:enhance the feature file \n\n

roost_feedback [08/07/2025, 12:45:20 PM]:enhance the feature file \n

roost_feedback [08/07/2025, 12:48:43 PM]:enhance the feature file\n\n

roost_feedback [08/07/2025, 12:55:42 PM]:remove unnecessary scenarios if any\n
*/

// ********RoostGPT********

Feature: YouTube API Testing

Background:
  Given the API base URL is "https://www.googleapis.com/youtube/v3"
  And the authorization header is set with a valid API key
  And the content type is "application/json"
  And the request timeout is set to 30 seconds
  And the response format is set to JSON
  And test data is initialized from the test data repository

Scenario: Search for videos by keyword
  When I send a GET request to "/search" with the following parameters:
    | part             | snippet       |
    | q                | test automation |
    | maxResults       | 5             |
    | type             | video         |
    | relevanceLanguage | en           |
  Then the response status code should be 200
  And the response should contain a list of videos
  And each video in the response should have an ID and title
  And the response time should be less than 2 seconds
  And the response should contain exactly 5 items
  And the response should include a pageInfo object with totalResults and resultsPerPage

Scenario: Get video details by ID
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet,statistics,contentDetails |
    | id         | dQw4w9WgXcQ                      |
  Then the response status code should be 200
  And the response should contain video details
  And the video details should include title, description, and view count
  And the video duration should be in ISO 8601 format
  And the response should include the channel information
  And the video statistics should include like count and comment count

Scenario: Handle invalid API key
  Given the authorization header is set with an invalid API key "AIzaSyInvalidKeyExample123456789"
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
  Then the response status code should be 403
  And the response should contain an error message about invalid credentials
  And the error reason should be "API key not valid"

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

Scenario: Get channel information
  When I send a GET request to "/channels" with the following parameters:
    | part       | snippet,statistics,contentDetails |
    | id         | UC_x5XG1OV2P6uZZ5FSM9Ttw         |
  Then the response status code should be 200
  And the response should contain channel details
  And the channel details should include title, description, and subscriber count
  And the channel statistics should include view count and video count

Scenario: Handle resource not found
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet       |
    | id         | nonExistentVideoId123 |
  Then the response status code should be 200
  And the response should contain an empty items array
  And the pageInfo.totalResults should be 0

Scenario: Verify video comments
  When I send a GET request to "/commentThreads" with the following parameters:
    | part       | snippet,replies |
    | videoId    | dQw4w9WgXcQ     |
    | maxResults | 10              |
  Then the response status code should be 200
  And the response should contain a list of comment threads
  And each comment should have an author, text, and publication date
  And top-level comments with replies should include reply data

Scenario: Test API error handling for invalid parameters
  When I send a GET request to "/search" with the following parameters:
    | part       | invalidPart   |
    | q          | test automation |
  Then the response status code should be 400
  And the response should contain a detailed error message
  And the error should indicate which parameter is invalid