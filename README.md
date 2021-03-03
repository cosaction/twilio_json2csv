Usage
=====
The intent of this tool is to be a quickly made application for converting json-based
Twilio log files into CSV for use by others in their spreadsheet of choice.

This is a command line application and as such must be run from the command line. Windows instructions
are given here as most users of this tool operate in Windows.

If this command is run by itself with no parameters, it will look in its current directory
for a file called 'twilio.json' as an input and output the data from that file to 'twilio.csv'
by default.

If you want to override the input (JSON) filename, you can do this by specifying '-in [fileName]'.

If you want to override the output (CSV) filename, you can do this by specifying '-out [fileName]'.

Usage Examples:

twilio_json2csv.exe                      -- This will read in twilio.json and output twilio.csv
                                     to/from the same directory that json2csv.exe is in

twilio_json2csv.exe -in otherFile.json   -- This will read in otherFile.sjon and output twilio.csv
                                     to/from the same directory that json2csv.exe is in

twilio_json2csv.exe -out newOutFile.csv  -- This will read in twilio.csv and output newOutFile.csv
                                     to/from the same directory that json2csv.exe in in


You may mix and match the -in and -out commands, as well as provide full paths to files to read
or write, for example:

twilio_json2csv.exe -in otherFile.json -out C:\Documents\newOutFile.csv

Would read-in otherFile.json and output the results to newOutFile.csv in the C:\Documents\ folder.
