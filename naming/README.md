# Naming: Maximum Obfuscation

Single letters, lying names, permanent typos, and variables that shadow built-ins. Every naming anti-pattern.

## What's Inside

**a()** - Function and all parameters are single letters. `a(b, c, d, e, f)` tells you nothing. Reading the code is decryption.

**data1 through data8** - Sequential numbering instead of descriptive names. Is `data3` a slice? Map? String? Who knows.

**strProcess** - Hungarian notation in Go. `strInput`, `intCount`, `boolFlag`. The type system already knows the types.

**UsrMgr** - Maximum abbreviation. `usrs`, `cfg`, `db`, `ctx`, `req`, `res`. Vowels are for the weak.

**ProcessUserDataFromDatabaseAndTransformItIntoJSONFormatAndSendItToTheClientViaHTTPResponse** - Function name that doesn't fit on one line. Need horizontal scroll to read it.

**Process_User_Data** - Mixes camelCase, snake_case, PascalCase in one function. `userName`, `user_id`, `UserEmail`. Consistency is boring.

**GetUser()** - Deletes all users. `DeleteCache()` creates a user. `IsValid()` returns int. Names lie about behavior.

**CalcualteTotal** - Permanent typo. `toatl`, `intput`, `otput`. Can't fix without breaking API. Typos are features.

**ShadowBuiltins** - Redefines `int`, `string`, `bool`, `error` as variables. "cannot use string as string value" compiler errors.

**DoStuff** - Generic meaningless names. `thing1`, `thing2`, `tmp`, `temp`, `stuff`, `result`. Intent is a mystery.

**MAGIC_NUMBER_1 = 86400** - Constants with no explanation. What does 86400 mean? Nobody documented it.

**UserUser** - Redundant prefixes everywhere. `UserName`, `UserEmail`, `UserPassword` in the `UserUser` struct.

**qwerty()** - Function named by keyboard mashing. Parameters `asdf` and `hjkl`. Variables `zxcv` and `qwer`.

## Why This Makes Code Unmaintainable

1. **Impossible to understand** - Code requires full context to comprehend
2. **Can't search** - Generic names appear everywhere
3. **No intent** - Names don't reveal what code does
4. **Permanent typos** - Can't fix without breaking compatibility
5. **Shadow bugs** - Shadowing built-ins causes bizarre compiler errors

## What You Should Do Instead

- Use descriptive, intention-revealing names
- Prefer clarity over brevity
- Use camelCase for unexported, PascalCase for exported
- Avoid abbreviations unless universally understood
- Never shadow built-in types
- Name by purpose, not implementation
- Use meaningful constant names
- Keep names pronounceable and searchable

## Do Not

Use these patterns unless you want nobody to understand your code, including yourself in 3 months.
