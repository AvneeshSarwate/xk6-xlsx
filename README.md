A module for https://k6.io/ to read excel files. Constructed in the pattern of https://k6.io/blog/extending-k6-with-xk6 and using excel-parsing capabilities of https://github.com/360EntSecGroup-Skylar/excelize.

Build agains the k6 binary using the command 

```
{xk6_binary} build v0.29.0 --with github.com/AvneeshSarwate/xk6-xlsx='/absolute/path/to/local/xk6-xlsx/directory'
```