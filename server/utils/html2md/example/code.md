# 7.13. Date and Time Functions and Operators

 
 

## Date and Time Operators

 

<table> <colgroup> <col width="9%"/> <col width="60%"/> <col width="31%"/> </colgroup> <thead valign="bottom"> <tr><th>Operator</th> <th>Example</th> <th>Result</th> </tr> </thead> <tbody valign="top"> <tr><td><code>\+</code></td> <td><code>date &#39;2012\-08\-08&#39; \+ interval &#39;2&#39; day</code></td> <td><code>2012\-08\-10</code></td> </tr> <tr><td><code>\+</code></td> <td><code>time &#39;01:00&#39; \+ interval &#39;3&#39; hour</code></td> <td><code>04:00:00.000</code></td> </tr> <tr><td><code>\+</code></td> <td><code>timestamp &#39;2012\-08\-08 01:00&#39; \+ interval &#39;29&#39; hour</code></td> <td><code>2012\-08\-09 06:00:00.000</code></td> </tr> <tr><td><code>\+</code></td> <td><code>timestamp &#39;2012\-10\-31 01:00&#39; \+ interval &#39;1&#39; month</code></td> <td><code>2012\-11\-30 01:00:00.000</code></td> </tr> <tr><td><code>\+</code></td> <td><code>interval &#39;2&#39; day \+ interval &#39;3&#39; hour</code></td> <td><code>2 03:00:00.000</code></td> </tr> <tr><td><code>\+</code></td> <td><code>interval &#39;3&#39; year \+ interval &#39;5&#39; month</code></td> <td><code>3\-5</code></td> </tr> <tr><td><code>\-</code></td> <td><code>date &#39;2012\-08\-08&#39; \- interval &#39;2&#39; day</code></td> <td><code>2012\-08\-06</code></td> </tr> <tr><td><code>\-</code></td> <td><code>time &#39;01:00&#39; \- interval &#39;3&#39; hour</code></td> <td><code>22:00:00.000</code></td> </tr> <tr><td><code>\-</code></td> <td><code>timestamp &#39;2012\-08\-08 01:00&#39; \- interval &#39;29&#39; hour</code></td> <td><code>2012\-08\-06 20:00:00.000</code></td> </tr> <tr><td><code>\-</code></td> <td><code>timestamp &#39;2012\-10\-31 01:00&#39; \- interval &#39;1&#39; month</code></td> <td><code>2012\-09\-30 01:00:00.000</code></td> </tr> <tr><td><code>\-</code></td> <td><code>interval &#39;2&#39; day \- interval &#39;3&#39; hour</code></td> <td><code>1 21:00:00.000</code></td> </tr> <tr><td><code>\-</code></td> <td><code>interval &#39;3&#39; year \- interval &#39;5&#39; month</code></td> <td><code>2\-7</code></td> </tr> </tbody> </table>

 
 
 

## Time Zone Conversion

 
The <code>AT TIME ZONE</code> operator sets the time zone of a timestamp:
 



<pre>SELECT timestamp &#39;2012-10-31 01:00 UTC&#39;;
2012-10-31 01:00:00.000 UTC

SELECT timestamp &#39;2012-10-31 01:00 UTC&#39; AT TIME ZONE &#39;America/Los_Angeles&#39;;
2012-10-30 18:00:00.000 America/Los_Angeles
</pre>

 
 
 
 

## Date and Time Functions

 

 
-  <code>current\_date \-&gt; date</code> 
- 
Returns the current date as of the start of the query.
 

 

 
-  <code>current\_time \-&gt; time with time zone</code> 
- 
Returns the current time as of the start of the query.
 

 

 
-  <code>current\_timestamp \-&gt; timestamp with time zone</code> 
- 
Returns the current timestamp as of the start of the query.
 

 

 
-  <code>current\_timezone</code>() → varchar 
- 
Returns the current time zone in the format defined by IANA (e.g., <code>America/Los\_Angeles</code>) or as fixed offset from UTC (e.g., <code>\+08:35</code>)
 

 

 
-  <code>date</code>(x) → date 
- 
This is an alias for <code>CAST(x AS date)</code>.
 

 

 
-  <code>from\_iso8601\_timestamp</code>(string) → timestamp with time zone 
- 
Parses the ISO 8601 formatted <code>string</code> into a <code>timestamp with time zone</code>.
 

 

 
-  <code>from\_iso8601\_date</code>(string) → date 
- 
Parses the ISO 8601 formatted <code>string</code> into a <code>date</code>.
 

 

 
-  <code>from\_unixtime</code>(unixtime) → timestamp 
- 
Returns the UNIX timestamp <code>unixtime</code> as a timestamp.
 

 

 
-  <code>from\_unixtime</code>(unixtime, string) → timestamp with time zone 
- 
Returns the UNIX timestamp <code>unixtime</code> as a timestamp with time zone using <code>string</code> for the time zone.
 

 

 
-  <code>from\_unixtime</code>(unixtime, hours, minutes) → timestamp with time zone 
- 
Returns the UNIX timestamp <code>unixtime</code> as a timestamp with time zone using <code>hours</code> and <code>minutes</code> for the time zone offset.
 

 

 
-  <code>localtime \-&gt; time</code> 
- 
Returns the current time as of the start of the query.
 

 

 
-  <code>localtimestamp \-&gt; timestamp</code> 
- 
Returns the current timestamp as of the start of the query.
 

 

 
-  <code>now</code>() → timestamp with time zone 
- 
This is an alias for <code>current\_timestamp</code>.
 

 

 
-  <code>to\_iso8601</code>(x) → varchar 
- 
Formats <code>x</code> as an ISO 8601 string. <code>x</code> can be date, timestamp, or timestamp with time zone.
 

 

 
-  <code>to\_milliseconds</code>(interval) → bigint 
- 
Returns the day-to-second <code>interval</code> as milliseconds.
 

 

 
-  <code>to\_unixtime</code>(timestamp) → double 
- 
Returns <code>timestamp</code> as a UNIX timestamp.
 

 
 
Note
 
The following SQL-standard functions do not use parenthesis:
 

 
- <code>current\_date</code> 
- <code>current\_time</code> 
- <code>current\_timestamp</code> 
- <code>localtime</code> 
- <code>localtimestamp</code> 

 
 
 
 

## Truncation Function

 
The <code>date\_trunc</code> function supports the following units:
 

<table> <colgroup> <col width="29%"/> <col width="71%"/> </colgroup> <thead valign="bottom"> <tr><th>Unit</th> <th>Example Truncated Value</th> </tr> </thead> <tbody valign="top"> <tr><td><code>second</code></td> <td><code>2001\-08\-22 03:04:05.000</code></td> </tr> <tr><td><code>minute</code></td> <td><code>2001\-08\-22 03:04:00.000</code></td> </tr> <tr><td><code>hour</code></td> <td><code>2001\-08\-22 03:00:00.000</code></td> </tr> <tr><td><code>day</code></td> <td><code>2001\-08\-22 00:00:00.000</code></td> </tr> <tr><td><code>week</code></td> <td><code>2001\-08\-20 00:00:00.000</code></td> </tr> <tr><td><code>month</code></td> <td><code>2001\-08\-01 00:00:00.000</code></td> </tr> <tr><td><code>quarter</code></td> <td><code>2001\-07\-01 00:00:00.000</code></td> </tr> <tr><td><code>year</code></td> <td><code>2001\-01\-01 00:00:00.000</code></td> </tr> </tbody> </table>

 
The above examples use the timestamp <code>2001\-08\-22 03:04:05.321</code> as the input.
 

 
-  <code>date\_trunc</code>(unit, x) → [same as input] 
- 
Returns <code>x</code> truncated to <code>unit</code>.
 

 
 
 

## Interval Functions

 
The functions in this section support the following interval units:
 

<table> <colgroup> <col width="49%"/> <col width="51%"/> </colgroup> <thead valign="bottom"> <tr><th>Unit</th> <th>Description</th> </tr> </thead> <tbody valign="top"> <tr><td><code>millisecond</code></td> <td>Milliseconds</td> </tr> <tr><td><code>second</code></td> <td>Seconds</td> </tr> <tr><td><code>minute</code></td> <td>Minutes</td> </tr> <tr><td><code>hour</code></td> <td>Hours</td> </tr> <tr><td><code>day</code></td> <td>Days</td> </tr> <tr><td><code>week</code></td> <td>Weeks</td> </tr> <tr><td><code>month</code></td> <td>Months</td> </tr> <tr><td><code>quarter</code></td> <td>Quarters of a year</td> </tr> <tr><td><code>year</code></td> <td>Years</td> </tr> </tbody> </table>

 

 
-  <code>date\_add</code>(unit, value, timestamp) → [same as input] 
- 
Adds an interval <code>value</code> of type <code>unit</code> to <code>timestamp</code>. Subtraction can be performed by using a negative value.
 

 

 
-  <code>date\_diff</code>(unit, timestamp1, timestamp2) → bigint 
- 
Returns <code>timestamp2 \- timestamp1</code> expressed in terms of <code>unit</code>.
 

 
 
 

## Duration Function

 
The <code>parse\_duration</code> function supports the following units:
 

<table> <colgroup> <col width="35%"/> <col width="65%"/> </colgroup> <thead valign="bottom"> <tr><th>Unit</th> <th>Description</th> </tr> </thead> <tbody valign="top"> <tr><td><code>ns</code></td> <td>Nanoseconds</td> </tr> <tr><td><code>us</code></td> <td>Microseconds</td> </tr> <tr><td><code>ms</code></td> <td>Milliseconds</td> </tr> <tr><td><code>s</code></td> <td>Seconds</td> </tr> <tr><td><code>m</code></td> <td>Minutes</td> </tr> <tr><td><code>h</code></td> <td>Hours</td> </tr> <tr><td><code>d</code></td> <td>Days</td> </tr> </tbody> </table>

 

 
-  <code>parse\_duration</code>(string) → interval 
- 
Parses <code>string</code> of format <code>value unit</code> into an interval, where <code>value</code> is fractional number of <code>unit</code> values:
 



<pre>SELECT parse_duration(&#39;42.8ms&#39;); -- 0 00:00:00.043
SELECT parse_duration(&#39;3.81 d&#39;); -- 3 19:26:24.000
SELECT parse_duration(&#39;5m&#39;);     -- 0 00:05:00.000
</pre>

 
 

 
 
 

## MySQL Date Functions

 
The functions in this section use a format string that is compatible with the MySQL <code>date\_parse</code> and <code>str\_to\_date</code> functions. The following table, based on the MySQL manual, describes the format specifiers:
 

<table> <colgroup> <col width="7%"/> <col width="93%"/> </colgroup> <thead valign="bottom"> <tr><th>Specifier</th> <th>Description</th> </tr> </thead> <tbody valign="top"> <tr><td><code>%a</code></td> <td>Abbreviated weekday name (<code>Sun</code> .. <code>Sat</code>)</td> </tr> <tr><td><code>%b</code></td> <td>Abbreviated month name (<code>Jan</code> .. <code>Dec</code>)</td> </tr> <tr><td><code>%c</code></td> <td>Month, numeric (<code>1</code> .. <code>12</code>) [[4]](#z)</td> </tr> <tr><td><code>%D</code></td> <td>Day of the month with English suffix (<code>0th</code>, <code>1st</code>, <code>2nd</code>, <code>3rd</code>, …)</td> </tr> <tr><td><code>%d</code></td> <td>Day of the month, numeric (<code>01</code> .. <code>31</code>) [[4]](#z)</td> </tr> <tr><td><code>%e</code></td> <td>Day of the month, numeric (<code>1</code> .. <code>31</code>) [[4]](#z)</td> </tr> <tr><td><code>%f</code></td> <td>Fraction of second (6 digits for printing: <code>000000</code> .. <code>999000</code>; 1 - 9 digits for parsing: <code>0</code> .. <code>999999999</code>) [[1]](#f)</td> </tr> <tr><td><code>%H</code></td> <td>Hour (<code>00</code> .. <code>23</code>)</td> </tr> <tr><td><code>%h</code></td> <td>Hour (<code>01</code> .. <code>12</code>)</td> </tr> <tr><td><code>%I</code></td> <td>Hour (<code>01</code> .. <code>12</code>)</td> </tr> <tr><td><code>%i</code></td> <td>Minutes, numeric (<code>00</code> .. <code>59</code>)</td> </tr> <tr><td><code>%j</code></td> <td>Day of year (<code>001</code> .. <code>366</code>)</td> </tr> <tr><td><code>%k</code></td> <td>Hour (<code>0</code> .. <code>23</code>)</td> </tr> <tr><td><code>%l</code></td> <td>Hour (<code>1</code> .. <code>12</code>)</td> </tr> <tr><td><code>%M</code></td> <td>Month name (<code>January</code> .. <code>December</code>)</td> </tr> <tr><td><code>%m</code></td> <td>Month, numeric (<code>01</code> .. <code>12</code>) [[4]](#z)</td> </tr> <tr><td><code>%p</code></td> <td><code>AM</code> or <code>PM</code></td> </tr> <tr><td><code>%r</code></td> <td>Time, 12-hour (<code>hh:mm:ss</code> followed by <code>AM</code> or <code>PM</code>)</td> </tr> <tr><td><code>%S</code></td> <td>Seconds (<code>00</code> .. <code>59</code>)</td> </tr> <tr><td><code>%s</code></td> <td>Seconds (<code>00</code> .. <code>59</code>)</td> </tr> <tr><td><code>%T</code></td> <td>Time, 24-hour (<code>hh:mm:ss</code>)</td> </tr> <tr><td><code>%U</code></td> <td>Week (<code>00</code> .. <code>53</code>), where Sunday is the first day of the week</td> </tr> <tr><td><code>%u</code></td> <td>Week (<code>00</code> .. <code>53</code>), where Monday is the first day of the week</td> </tr> <tr><td><code>%V</code></td> <td>Week (<code>01</code> .. <code>53</code>), where Sunday is the first day of the week; used with <code>%X</code></td> </tr> <tr><td><code>%v</code></td> <td>Week (<code>01</code> .. <code>53</code>), where Monday is the first day of the week; used with <code>%x</code></td> </tr> <tr><td><code>%W</code></td> <td>Weekday name (<code>Sunday</code> .. <code>Saturday</code>)</td> </tr> <tr><td><code>%w</code></td> <td>Day of the week (<code>0</code> .. <code>6</code>), where Sunday is the first day of the week [[3]](#w)</td> </tr> <tr><td><code>%X</code></td> <td>Year for the week where Sunday is the first day of the week, numeric, four digits; used with <code>%V</code></td> </tr> <tr><td><code>%x</code></td> <td>Year for the week, where Monday is the first day of the week, numeric, four digits; used with <code>%v</code></td> </tr> <tr><td><code>%Y</code></td> <td>Year, numeric, four digits</td> </tr> <tr><td><code>%y</code></td> <td>Year, numeric (two digits) [[2]](#y)</td> </tr> <tr><td><code>%%</code></td> <td>A literal <code>%</code> character</td> </tr> <tr><td><code>%x</code></td> <td><code>x</code>, for any <code>x</code> not listed above</td> </tr> </tbody> </table>

 

<table rules="none"> <colgroup><col class="label"/><col/></colgroup> <tbody valign="top"> <tr><td>[[1]](#id4)</td><td>Timestamp is truncated to milliseconds.</td></tr> </tbody> </table>

 

<table rules="none"> <colgroup><col class="label"/><col/></colgroup> <tbody valign="top"> <tr><td>[[2]](#id7)</td><td>When parsing, two-digit year format assumes range <code>1970</code> .. <code>2069</code>, so “70” will result in year <code>1970</code> but “69” will produce <code>2069</code>.</td></tr> </tbody> </table>

 

<table rules="none"> <colgroup><col class="label"/><col/></colgroup> <tbody valign="top"> <tr><td>[[3]](#id6)</td><td>This specifier is not supported yet. Consider using [<code>day\_of\_week()</code>](#day_of_week) (it uses <code>1\-7</code> instead of <code>0\-6</code>).</td></tr> </tbody> </table>

 

<table rules="none"> <colgroup><col class="label"/><col/></colgroup> <tbody valign="top"> <tr><td>[4]</td><td><em>([1](#id1), [2](#id2), [3](#id3), [4](#id5))</em> This specifier does not support <code>0</code> as a month or day.</td></tr> </tbody> </table>

 
 
Warning
 
The following specifiers are not currently supported: <code>%D %U %u %V %w %X</code>
 
 

 
-  <code>date\_format</code>(timestamp, format) → varchar 
- 
Formats <code>timestamp</code> as a string using <code>format</code>.
 

 

 
-  <code>date\_parse</code>(string, format) → timestamp 
- 
Parses <code>string</code> into a timestamp using <code>format</code>.
 

 
 
 

## Java Date Functions

 
The functions in this section use a format string that is compatible with JodaTime’s [DateTimeFormat](http://joda-time.sourceforge.net/apidocs/org/joda/time/format/DateTimeFormat.html) pattern format.
 

 
-  <code>format\_datetime</code>(timestamp, format) → varchar 
- 
Formats <code>timestamp</code> as a string using <code>format</code>.
 

 

 
-  <code>parse\_datetime</code>(string, format) → timestamp with time zone 
- 
Parses <code>string</code> into a timestamp with time zone using <code>format</code>.
 

 
 
 

## Extraction Function

 
The <code>extract</code> function supports the following fields:
 

<table> <colgroup> <col width="45%"/> <col width="55%"/> </colgroup> <thead valign="bottom"> <tr><th>Field</th> <th>Description</th> </tr> </thead> <tbody valign="top"> <tr><td><code>YEAR</code></td> <td>[<code>year()</code>](#year)</td> </tr> <tr><td><code>QUARTER</code></td> <td>[<code>quarter()</code>](#quarter)</td> </tr> <tr><td><code>MONTH</code></td> <td>[<code>month()</code>](#month)</td> </tr> <tr><td><code>WEEK</code></td> <td>[<code>week()</code>](#week)</td> </tr> <tr><td><code>DAY</code></td> <td>[<code>day()</code>](#day)</td> </tr> <tr><td><code>DAY\_OF\_MONTH</code></td> <td>[<code>day()</code>](#day)</td> </tr> <tr><td><code>DAY\_OF\_WEEK</code></td> <td>[<code>day\_of\_week()</code>](#day_of_week)</td> </tr> <tr><td><code>DOW</code></td> <td>[<code>day\_of\_week()</code>](#day_of_week)</td> </tr> <tr><td><code>DAY\_OF\_YEAR</code></td> <td>[<code>day\_of\_year()</code>](#day_of_year)</td> </tr> <tr><td><code>DOY</code></td> <td>[<code>day\_of\_year()</code>](#day_of_year)</td> </tr> <tr><td><code>YEAR\_OF\_WEEK</code></td> <td>[<code>year\_of\_week()</code>](#year_of_week)</td> </tr> <tr><td><code>YOW</code></td> <td>[<code>year\_of\_week()</code>](#year_of_week)</td> </tr> <tr><td><code>HOUR</code></td> <td>[<code>hour()</code>](#hour)</td> </tr> <tr><td><code>MINUTE</code></td> <td>[<code>minute()</code>](#minute)</td> </tr> <tr><td><code>SECOND</code></td> <td>[<code>second()</code>](#second)</td> </tr> <tr><td><code>TIMEZONE\_HOUR</code></td> <td>[<code>timezone\_hour()</code>](#timezone_hour)</td> </tr> <tr><td><code>TIMEZONE\_MINUTE</code></td> <td>[<code>timezone\_minute()</code>](#timezone_minute)</td> </tr> </tbody> </table>

 
The types supported by the <code>extract</code> function vary depending on the field to be extracted. Most fields support all date and time types.
 

 
-  <code>extract</code>(field FROM x) → bigint 
- 
Returns <code>field</code> from <code>x</code>.
 
 
Note
 
This SQL-standard function uses special syntax for specifying the arguments.
 
 

 
 
 

## Convenience Extraction Functions

 

 
-  <code>day</code>(x) → bigint 
- 
Returns the day of the month from <code>x</code>.
 

 

 
-  <code>day\_of\_month</code>(x) → bigint 
- 
This is an alias for [<code>day()</code>](#day).
 

 

 
-  <code>day\_of\_week</code>(x) → bigint 
- 
Returns the ISO day of the week from <code>x</code>. The value ranges from <code>1</code> (Monday) to <code>7</code> (Sunday).
 

 

 
-  <code>day\_of\_year</code>(x) → bigint 
- 
Returns the day of the year from <code>x</code>. The value ranges from <code>1</code> to <code>366</code>.
 

 

 
-  <code>dow</code>(x) → bigint 
- 
This is an alias for [<code>day\_of\_week()</code>](#day_of_week).
 

 

 
-  <code>doy</code>(x) → bigint 
- 
This is an alias for [<code>day\_of\_year()</code>](#day_of_year).
 

 

 
-  <code>hour</code>(x) → bigint 
- 
Returns the hour of the day from <code>x</code>. The value ranges from <code>0</code> to <code>23</code>.
 

 

 
-  <code>millisecond</code>(x) → bigint 
- 
Returns the millisecond of the second from <code>x</code>.
 

 

 
-  <code>minute</code>(x) → bigint 
- 
Returns the minute of the hour from <code>x</code>.
 

 

 
-  <code>month</code>(x) → bigint 
- 
Returns the month of the year from <code>x</code>.
 

 

 
-  <code>quarter</code>(x) → bigint 
- 
Returns the quarter of the year from <code>x</code>. The value ranges from <code>1</code> to <code>4</code>.
 

 

 
-  <code>second</code>(x) → bigint 
- 
Returns the second of the minute from <code>x</code>.
 

 

 
-  <code>timezone\_hour</code>(timestamp) → bigint 
- 
Returns the hour of the time zone offset from <code>timestamp</code>.
 

 

 
-  <code>timezone\_minute</code>(timestamp) → bigint 
- 
Returns the minute of the time zone offset from <code>timestamp</code>.
 

 

 
-  <code>week</code>(x) → bigint 
- 
Returns the [ISO week](https://en.wikipedia.org/wiki/ISO_week_date) of the year from <code>x</code>. The value ranges from <code>1</code> to <code>53</code>.
 

 

 
-  <code>week\_of\_year</code>(x) → bigint 
- 
This is an alias for [<code>week()</code>](#week).
 

 

 
-  <code>year</code>(x) → bigint 
- 
Returns the year from <code>x</code>.
 

 

 
-  <code>year\_of\_week</code>(x) → bigint 
- 
Returns the year of the [ISO week](https://en.wikipedia.org/wiki/ISO_week_date) from <code>x</code>.
 

 

 
-  <code>yow</code>(x) → bigint 
- 
This is an alias for [<code>year\_of\_week()</code>](#year_of_week).
 

 
 
 
 