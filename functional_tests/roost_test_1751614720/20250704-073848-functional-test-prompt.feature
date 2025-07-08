
// ********RoostGPT********
/*

roost_feedback [04/07/2025, 1:20:57 PM]:enhance and improve the test\n\n

roost_feedback [08/07/2025, 12:25:39 PM]:enhance the test feature file\n\n

roost_feedback [08/07/2025, 12:41:22 PM]:enhance the feature file \n\n

roost_feedback [08/07/2025, 12:45:20 PM]:enhance the feature file \n

roost_feedback [08/07/2025, 12:48:43 PM]:enhance the feature file\n\n
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
  And rate limiting parameters are configured appropriately
  And response validation schemas are loaded
  And test data is initialized from the test data repository
  And performance monitoring is enabled for all requests

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
  And the etag in the response should be a non-empty string
  And the response should include a pageInfo object with totalResults and resultsPerPage
  And all videos should be relevant to the search query "test automation"
  And the response should be saved for regression testing

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
  And the video's license information should be present
  And the contentDetails should include information about content rating
  And the video's tags should be returned as an array if present
  And the video's localized information should match the requested language

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
  And the response should include a detailed error code
  And the response content type should be "application/json"
  And the error should be logged with appropriate severity level
  And the system should recommend troubleshooting steps

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
  And the pageInfo.totalResults should be consistent across pagination requests
  And the regionCode in the response should match the expected region
  And pagination metadata should be consistent across all pages
  And the last page should not contain a nextPageToken

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
  And each video's publishedAt timestamp should be in ISO 8601 format
  And the response should include videos from various channels
  And the date filtering should work correctly with different timezones
  And the response should include videos from the entire date range
  And the date filtering should respect the exact time components

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
  And the channel's custom URL should be present if available
  And the channel's topic categories should be listed if available
  And the channel's creation date should be in ISO 8601 format
  And the channel's localization information should be present
  And the channel's content owner details should be included if applicable

Scenario: Handle resource not found
  When I send a GET request to "/videos" with the following parameters:
    | part       | snippet       |
    | id         | nonExistentVideoId123 |
  Then the response status code should be 200
  And the response should contain an empty items array
  And the response should include a valid kind property
  And the response should include a valid etag
  And the response should not contain any error messages
  And the pageInfo.totalResults should be 0
  And the response time should be less than 1 second
  And the system should log the not-found resource appropriately
  And the response structure should match successful responses
  And proper error handling should be demonstrated

Scenario: Rate limiting behavior
  When I send 10 consecutive GET requests to "/search" with minimal delay
  Then all responses should have status code 200 or 429
  And if status code is 429, the response should contain quota exceeded information
  And the response headers should include rate limiting information
  And the quota usage information should be present in the response
  And the error response should include a retry-after header if rate limited
  And the error message should provide guidance on quota management
  And the error response should include a quotaExceeded error code if applicable
  And the response should include information about daily quota limits
  And the system should implement exponential backoff for retries
  And quota consumption should be tracked and reported

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
  And the response should include category-specific metadata
  And the video snippets should contain relevant category information
  And the category name should be consistent with the category ID
  And the videos should be properly categorized according to YouTube guidelines
  And the category filtering should work in conjunction with other filters

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
  And the related videos should share similar tags or categories
  And the response should include relevance information for each video
  And the related videos should have similar content characteristics
  And the algorithm should prioritize videos from the same channel
  And the related videos should have similar audience demographics
  And the relevance scoring should be consistent with YouTube's web interface

Scenario: Search for live streams
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | eventType  | live          |
    | maxResults | 5             |
    | type       | video         |
  Then the response status code should be 200
  And all videos in the response should have a live broadcast content status
  And the response should include live streaming details for each video
  And each live stream should have a concurrent viewer count if available
  And the live streams should include scheduled start times
  And the response should indicate if the stream is currently active
  And the live status should be accurately reflected in the response
  And upcoming streams should be distinguished from currently active streams
  And completed live streams should not be included in the results
  And the response should include stream health metrics if available

Scenario: Verify video comments
  When I send a GET request to "/commentThreads" with the following parameters:
    | part       | snippet,replies |
    | videoId    | dQw4w9WgXcQ     |
    | maxResults | 10              |
  Then the response status code should be 200
  And the response should contain a list of comment threads
  And each comment should have an author, text, and publication date
  And the comments should be sorted by relevance by default
  And the response should include reply counts for each comment thread
  And top-level comments with replies should include reply data
  And comment timestamps should be in ISO 8601 format
  And comment moderation status should be indicated if applicable
  And comment like counts should be included in the response
  And the comment author details should include channel information
  And comments should be properly formatted with HTML entities escaped

Scenario: Search for videos with specific duration
  When I send a GET request to "/search" with the following parameters:
    | part       | snippet       |
    | q          | test automation |
    | maxResults | 5             |
    | type       | video         |
    | videoDuration | medium     |
  Then the response status code should be 200
  And all videos in the response should have a medium duration
  And the response should contain a list of videos
  And each video should have duration information in the contentDetails
  And the videos should match the search query criteria
  And the duration filter should correctly classify videos
  And the exact duration in seconds should be available in the response
  And the duration filter should work in conjunction with other filters
  And the response should include videos from the entire duration range

Scenario: Verify video captions availability
  When I send a GET request to "/captions" with the following parameters:
    | part       | snippet       |
    | videoId    | dQw4w9WgXcQ   |
  Then the response status code should be 200
  And the response should indicate if captions are available
  And caption language information should be included
  And auto-generated captions should be distinguished from manual captions
  And caption formats should be specified in the response
  And caption track details should include timing information
  And the response should include caption track URLs if accessible

Scenario: Test API error handling for invalid parameters
  When I send a GET request to "/search" with the following parameters:
    | part       | invalidPart   |
    | q          | test automation |
  Then the response status code should be 400
  And the response should contain a detailed error message
  And the error should indicate which parameter is invalid
  And the response should include suggestions for valid parameters
  And the error should be properly formatted according to the API error schema
  And the error response should include a request ID for troubleshooting