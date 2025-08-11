Feature: RoostGPT CLI Functional and Non-Functional Testing

Background:
  Given the API base URL 'http://localhost:3000'
  And the authorization header is set
  And the content type is 'application/json'

Scenario: Verify Git Config Modification Prevention
  Given the ~/.git-credentials file is removed
  And multiple .env files with different git credentials are prepared
  When I execute RoostGPT CLI using a prepared .env file
  Then the system's GIT configuration should remain unchanged
  And 'user.name' and 'user.email' should not be modified

Scenario: Environment-Specific Credential Handling
  Given a .env file with specific git credentials is prepared
  When I execute RoostGPT CLI using the prepared .env file
  And perform a git operation using RoostGPT CLI
  Then the operation should use the credentials from the .env file
  And the system's GIT configuration should remain unchanged

Scenario: Performance Under Multiple Environment Files
  Given multiple .env files with different git credentials are prepared
  When I execute RoostGPT CLI with the first .env file
  Then I should record the time taken and resource usage
  When I repeat the process with additional .env files
  Then the performance metrics should show minimal time variation and resource usage

Scenario: Security of Credential Handling
  Given a .env file with dummy sensitive credentials is prepared
  When I execute RoostGPT CLI with the .env file
  Then no credentials from the .env file should be visible in logs, output, or any temporary files
  And system logs and temporary files should not show any credential leakage
