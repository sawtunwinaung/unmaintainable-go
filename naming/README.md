# Naming Conventions: The Art of Obfuscation

Welcome to the masterclass in variable naming! This directory showcases cutting-edge techniques for making your code as unreadable as possible through strategic naming choices.

## File: `naming_disaster.go`

### Revolutionary Techniques Demonstrated

#### 1. **Single Letter Variables Everywhere**
Function `a` demonstrates the ultimate in conciseness:
- Why use descriptive names when you have the entire alphabet?
- `a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`, `i` - so organized!
- Reading code becomes a fun alphabet puzzle
- Variable scope? Just remember which letter does what!

**Why it's genius**: Less typing = more productivity!

#### 2. **The data1, data2, data3 Pattern**
Sequential numbering is peak organization:
- `data1` through `data8` for all your data needs
- Easy to remember: just increment the number!
- No need to think about what the variable represents
- Perfect for job security through obscurity

**Why it's genius**: Consistent naming scheme!

#### 3. **Hungarian Notation in Go**
`strProcess` shows that C conventions never die:
- `strInput`, `intCount`, `boolFlag` - so descriptive!
- The type system already knows the types, but let's repeat it!
- Go's type inference? We don't trust it
- Makes refactoring extra fun when types change

**Why it's genius**: The variable name tells you the type twice!

#### 4. **Maximal Abbreviation**
`UsrMgr` and friends embrace the lost art of abbreviation:
- Why write `users` when `usrs` saves 1 character?
- `cfg`, `db`, `ctx`, `req`, `res` - everyone knows what these mean!
- Vowels are optional in professional code
- `PrcUsr` is obviously "process user"

**Why it's genius**: So concise! Like texting from 2005!

#### 5. **Absurdly Long Names**
`ProcessUserDataFromDatabaseAndTransformItIntoJSONFormatAndSendItToTheClientViaHTTPResponse`:
- Function names that don't fit on one line
- Describes implementation instead of intent
- Need a scroll bar just to read the name
- Comments are unnecessary when the name is 20 words long

**Why it's genius**: Absolutely no ambiguity!

#### 6. **Inconsistent Naming Styles**
`Process_User_Data` mixes camelCase, snake_case, and PascalCase:
- `userName`, `user_id`, `UserEmail` in the same function!
- `local_variable`, `localVariable`, `Local_Variable` in the same scope!
- Keeps you on your toes
- Every variable is special and unique

**Why it's genius**: Variety is the spice of life!

#### 7. **Names That Lie**
`GetUser`, `DeleteCache`, `IsValid` do the opposite of what they say:
- `GetUser()` deletes all users - surprise!
- `DeleteCache()` creates a user - plot twist!
- `IsValid()` returns an int - boolean schmolean!
- Keeps code reviewers alert

**Why it's genius**: Trust no one, not even function names!

#### 8. **Permanent Typos**
`CalcualteTotal`, `ProccessData`, `intput`, `otput`, `toatl`:
- Typos become part of the API
- Can't fix them without breaking compatibility
- Hiring developers with dyslexia gives them an advantage
- Spell checkers are for the weak

**Why it's genius**: Unique identifiers guaranteed!

#### 9. **Shadowing Built-in Types**
`ShadowBuiltins` redefines `int`, `string`, `bool`, `error`:
- Why use built-in types when you can redefine them?
- `var int int = 5` is perfectly valid Go!
- Creates compiler errors that look like brain teasers
- "cannot use string (variable of type string) as string value"

**Why it's genius**: Keeps the compiler on its toes!

#### 10. **Generic Meaningless Names**
`DoStuff`, `HandleThing`, `thing1`, `thing2`, `tmp`, `temp`:
- `thing` is the ultimate variable name
- `stuff` captures the essence of any data
- `tmp` and `temp` - because one temporary isn't enough
- Intent? That's for documentation!

**Why it's genius**: One name fits all use cases!

#### 11. **Mystery Constants**
`MAGIC_NUMBER_1 = 86400`, `THE_ANSWER = 54`:
- What do these numbers mean? Nobody knows!
- 86400 is definitely not seconds in a day (it is)
- THE_ANSWER is 54, not 42 - we're rebels!
- Comments explaining constants are for amateurs

**Why it's genius**: Job security through magical numbers!

#### 12. **Redundant Prefixes**
`UserUser` struct with `UserName`, `UserEmail`, `UserPassword`:
- Every field reminds you it's part of User!
- `GetUserName()` returns the `UserName` of the `UserUser`
- Prevents confusion with other Names and Emails
- More User = more professional

**Why it's genius**: Absolutely no ambiguity! Very clear!

## Why This Code is "Best Practice"

1. **Conciseness**: Single letters save typing time
2. **Consistency**: Numbers 1-8 are very consistent
3. **Clarity**: 50-word function names are crystal clear
4. **Uniqueness**: Typos make your code one-of-a-kind
5. **Job Security**: Only you can maintain this

## Usage

```go
// Good luck understanding this
result := naming.a(1, 2, "test", true, []int{1, 2, 3})

// What does this process?
naming.strProcess("hello", 3, true)

// What does this do? The name lies!
naming.GetUser() // Deletes all users

// Single letter paradise
naming.DoStuff(naming.data1)
```

## Common Anti-Patterns Celebrated Here

1. **Single letter variables** - `a`, `b`, `c`, `x`, `y`, `z`
2. **Numbered variables** - `data1`, `data2`, `thing1`, `thing2`
3. **Abbreviations** - `usrs`, `cfg`, `ctx`, `prc`
4. **Hungarian notation** - `strValue`, `intCount`, `boolFlag`
5. **Overly long names** - Function names that span multiple lines
6. **Inconsistent styles** - Mixing camelCase, snake_case, PascalCase
7. **Lying names** - Functions that do the opposite of their name
8. **Typos** - Permanent misspellings in identifiers
9. **Shadowing** - Redefining built-in types and keywords
10. **Generic names** - `thing`, `stuff`, `data`, `tmp`
11. **Magic numbers** - Constants with no explanation
12. **Redundant prefixes** - `UserUserName` in the `User` struct

## WARNING

**This code demonstrates what NOT to do.** Proper naming in Go:
- Use descriptive, intention-revealing names
- Prefer clarity over brevity
- Use consistent naming style (camelCase for unexported, PascalCase for exported)
- Avoid abbreviations unless they're universally understood
- Don't shadow built-in types
- Name functions after what they do, not how they do it
- Use meaningful constant names
- Avoid redundant prefixes in struct fields
- Keep names pronounceable and searchable

But where's the mystery in that? 😈
