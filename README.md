# meta
* meta name(Directive)
* template
* patterns
* output

example
```
melon meta -d ./ \
--interface --sql --gorm
--handler='directives=sql:table,sql:select template=file:./template/some.tpl parrtern=* output_prefix=zz'
--handler='directives=sql:table,sql:select template=file:./template/some.tpl parrtern=* output_prefix=zz'
```