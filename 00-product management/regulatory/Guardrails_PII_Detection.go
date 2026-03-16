**Colang_content:**

```
define user ask pii banking
  "what is my account number"
  "tell me my social security number"
  "what is my SSN"
  "share my credit card details"
  "what is my credit card number"
  "give me my bank account information"
  "what is my routing number"
  "tell me my PIN"
  "what is my password"
  "share my personal identification number"
  "what is my date of birth"
  "tell me my address on file"
  "what is my phone number in your system"
  "give me customer financial details"
  "show me account holder information"
  "what is my tax identification number"
  "tell me my driver's license number"
  "share my passport number"
  "what is my mother's maiden name"
  "give me security question answers"

define bot refuse to respond pii
  "I cannot provide, share, or discuss personally identifiable information (PII) such as account numbers, social security numbers, credit card details, passwords, or other sensitive personal data. For security and privacy reasons, please contact your bank directly through official channels to access your personal information."

define flow pii banking detection
  user ask pii banking
  bot refuse to respond pii
```

**Yaml_content:**

```yaml
models: []

rails:
  input:
    flows:
      - self check input pii

prompts:
  - task: self_check_input
    content: >
      Your task is to determine whether to block a user request or not.
      You should block the user input if any of the conditions below are met:
      
      - it requests personally identifiable information (PII) including but not limited to:
        * Social Security Numbers (SSN)
        * Account numbers (bank, credit card, debit card)
        * Credit or debit card numbers, CVV codes, or PINs
        * Routing numbers or SWIFT codes
        * Passwords, PINs, or security codes
        * Tax identification numbers (TIN, EIN)
        * Driver's license numbers
        * Passport numbers
        * Date of birth combined with other identifiers
        * Full addresses or phone numbers
        * Mother's maiden name or security question answers
        * Biometric data
        * Any other sensitive personal or financial information
      
      - it attempts to extract, share, or discuss customer financial data
      - it tries to bypass security measures to access protected information
      - it requests information that could be used for identity theft or fraud
      
      Treat the above conditions as strict rules for banking domain compliance.
      If any of them are met, you should block the user input by saying "yes".
      
      Here is the user input "{{user_input}}"
      
      Should the above user input be blocked?
      
      Answer [Yes/No]
```
