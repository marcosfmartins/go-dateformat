# go-dateformat

this project is a study of the implementation of a date formatting lib.

| Spec |  Example  |                Description                |
|:----:|:---------:|:-----------------------------------------:|
|  %Y  |   2006    |             Full year 4 digits            |
|  %y  |    06     |               Year 2 digits               |
|  %m  |    01     |            Month number (01-12)           |
| %-m  |     1     |            Month number (1-12)            |
|  %B  |  January  |              Full month name              |
|  %b  |    Jan    |           Abbreviated month name          |
|  %A  |  Monday   |             Full weekday name             |
|  %a  |    Mon    |          Abbreviated weekday name         |
|  %d  |    02     |             Day number (01-31)            |
| %-d  |     2     |             Day number (1-31)             |
|  %j  |    002    |         Day of the year (001-366)         |
| %-j  |     2     |          Day of the year (1-366)          |
|  %U  |     1     |          week of the year (00–53)         |
|  %H  |    15     |            Hour number (00-23)            |
|  %l  |    03     |            Hour number (00-12)            |
| %-l  |     3     |             Hour number (0-12)            |
|  %M  |    04     |           Minute number (00–59)           |
| %-M  |     4     |            Minute number (0–59)           |
|  %S  |    05     |           Second number (00–60)           |
| %-S  |     5     |            Second number (0–60)           |
| %3f  |    000    |            Milliseconds number            |
| %6f  |  000000   |            MicroSeconds number            |
|  %f  | 000000000 |             Nanoseconds number            |
| %9f  | 000000000 |             Nanoseconds number            |
|  %P  |   AM/PM   |         AM or PM in 12-hour clocks        |
|  %p  |   am/pm   |         am or pm in 12-hour clocks        |
|  %z  |   -0300   |  Offset from the local time to UTC (HHMM) |
| %-z  |    -03    |   Offset from the local time to UTC (HH)  |
| %:z  |  -03:00   | Offset from the local time to UTC (HH:MM) |
