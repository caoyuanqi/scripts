#!/usr/bin/perl
# This perl script will copy any thing in the download into the kindle folder 'Book' or 'Now'
use strict;

my $directory = '/Users/Yuanqi/Downloads';
opendir (DIR, $directory) or die $!;

while (my $file = readdir(DIR))
{
    if ($file =~ ".*\.pdf")
    {
        if($ARGV[0] == "now")
        {
            print "Moving $file to kindle now\n";
            `cp /Users/Yuanqi/Downloads/$file /Volumes/Kindle/Now`;
            `rm /Users/Yuanqi/Downloads/$file`;
        }
        elsif ($ARGV[0] == "book")
        {
                        print "Moving $file to kindle Books\n";
                        `cp /Users/Yuanqi/Downloads/$file /Volumes/Kindle/Books`;
                        `rm /Users/Yuanqi/Downloads/$file`;
        }
    }
}
