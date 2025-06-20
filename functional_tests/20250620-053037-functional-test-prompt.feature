Feature: YouTube Homepage Functionality

  Scenario: Verify homepage loads successfully
    Given the user has an internet connection
    When the user sends a GET request to "https://www.youtube.com/"
    Then the response status code should be 200
    And the response body should contain elements for "header"
    And the response body should contain elements for "sidebar"
    And the response body should contain elements for "recommendedVideos"

  Scenario: Verify recommended videos are displayed on homepage
    Given the user has an internet connection
    When the user sends a GET request to "https://www.youtube.com/"
    Then the response status code should be 200
    And the response body should contain an array of "recommendedVideos"
    And each video in "recommendedVideos" should have properties:
      | thumbnail   |
      | title       |
      | channelName |
      | viewCount   |
      | uploadTime  |

  Scenario: Verify trending section is available
    Given the user has an internet connection
    When the user sends a GET request to "https://www.youtube.com/feed/trending"
    Then the response status code should be 200
    And the response body should contain an array of "trendingVideos"
    And each video in "trendingVideos" should have properties:
      | thumbnail   |
      | title       |
      | channelName |
      | viewCount   |
      | uploadTime  |

  Scenario: Verify shorts section is available
    Given the user has an internet connection
    When the user sends a GET request to "https://www.youtube.com/shorts"
    Then the response status code should be 200
    And the response body should contain an array of "shortsVideos"
    And each video in "shortsVideos" should have properties:
      | thumbnail   |
      | title       |
      | channelName |
      | viewCount   |
