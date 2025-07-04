
// ********RoostGPT********
/*

roost_feedback [04/07/2025, 1:20:57 PM]:enhance and improve the test\n\n
*/

// ********RoostGPT********

Feature: YouTube API Testing

Background:
  Given the API base URL is "https://www.googleapis.com/youtube/v3"
  And the authorization header is set with a valid API key
  And the content type is "application/json"
  And the request timeout is set to 30 seconds

Scenario: Search for videos by keyword
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 5             |
    | type       | video         |
  Then the response status code should be 200
  And the response should contain a list of videos
  And each video in the response should have an ID and title

Scenario: Get video details by ID
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet,statistics,contentDetails |
    | id         | dQw4w9WgXcQ                      |
  Then the response status code should be 200
  And the response should contain video details
  And the video details should include title, description, and view count

Scenario: Handle invalid API key
  Given the authorization header is set with an invalid API key
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
  Then the response status code should be 403
  And the response should contain an error message about invalid credentials

Scenario: Verify pagination of search results
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 10            |
    | type       | video         |
  Then the response status code should be 200
  And the response should contain a nextPageToken
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 10            |
    | type       | video         |
    | pageToken  | {nextPageToken} |
  Then the response status code should be 200
  And the response should contain different videos than the first page