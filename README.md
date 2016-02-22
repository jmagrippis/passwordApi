# Memorable Password API

This is a simple golang API backend implementation of the [Golang Memorable Password Generator][].

Just a couple of endpoints that return an array of memorable passwords in json form:

```
/generate/:amount
/generate/:amount/safe
```

The safe endpoint churns out passwords featuring mixed case, numbers and symbols, that should satisfy most web forms.

## Sample Output

### /generate/5
```json
[
  "shouting mermaids fear posh teddy-bears",
  "shouting unicorns fancy homely teddy-bears",
  "slapping mermaids love posh boxes",
  "jumping piranhas eat posh teddy-bears",
  "shouting piranhas bring homely diamonds"
]
```

### /generate/3/safe
```json
[
  "Cuddling6Mermaids6Aggravate6Homely6Buckets(",
  "Slapping6Mermaids6Love6Posh6Boxes(",
  "Slapping6Mermaids6Love6Homely6Buckets("
]
```

[Golang Memorable Password Generator]: https://github.com/jmagrippis/password "I am told the creator is a rather dashing fellow."