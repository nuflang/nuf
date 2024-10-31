# Syntax for Nuf

## User defined text

<text>

```nuf
"Hello, World!"
```

## User defined name

--<text>_<text>_<text>

```nuf
--page_title
```

## Built-in functions

- landmark(<built-in-keyword>)
- heading(<built-in-keyword>)
- interaction(<built-in-keyword>)
- list(<array>)
- custom(--<user-defined-name>)

## Built-in keywords

- header
- aside
- footer
- form
- main
- nav
- section
- search

- navigation
- action

## Hashes

{
    name: <user-defined-name>,
    <expression>: <expression>,
}

## Array

[
    <expression>,
    <expression>,
    <expression>,
]

## Built-in calls

<built-in-function>(<built-in-keyword>, <hash>);

## Infixes

- |> - sibling (not necessarily phisically next to each other in DOM tree)
- -> - child

(<expression> -> <expression>) -> <expression>;
<expression> |> <expression>;

```nuf
heading("Content categories") -> landmark(section);
```

```nuf
"Log in" -> interaction(navigation, {
	href: "/login",
});
```

```nuf
"Departments"
-> interaction(navigation, {
    name: --departments_link,
    url: "/government/organisations",
})
-> custom(--government_activity_departments_title);
```

```nuf
list([
    custom(--government_activity_departments_title),
    custom(--government_activity_news_title),
    custom(--government_activity_guidance_and_regulation_title),
], {
    name: --government-activity_list,
}) |> custom(--government_activity_title);
```
