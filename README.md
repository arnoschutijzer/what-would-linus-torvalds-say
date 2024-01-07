# What Would Linus Torvalds Say

ChatGPT4.0 has way better results or you can try fine-tuning the model using r/linusrants or some of the examples in `fixtures/fine-tuning`.

Add a `.github/workflows/wwlts` file with the following contents:
```yml
on:
  pull_request:
    branches: [ main ]

jobs:
  review:
    runs-on: ubuntu-latest
    name: Review my patch pls
    steps:
      - name: Review
        uses: algleymi/what-would-linus-torvalds-say@v1.0.0
        env:
          GITHUB_TOKEN: ${{ github.token }}
          OPENAI_TOKEN: ${{ secrets.OPENAI_TOKEN }}
          OPENAI_MODEL: "gpt-3.5-turbo" # optional, works with any model the token has access to
```

<img src="./.github/torvalds.JPG" alt="Linus Torvalds giving you the finger." />