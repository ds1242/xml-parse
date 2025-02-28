# Bookplate Wordpress Export Cleanup
Go Script that runs through the XML Wordpress output for the Bookplate exhibits.
Created a Go Struct off the output and then loop through each level of the output and 
add it to an Post struct.  Each Post struct is then written to a CSV row.