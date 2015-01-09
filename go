#! /usr/bin/perl
use Cwd;

@list = ("study","code","movie","9319","note","ielts");
#@file_list_to_display;
$mark = 0;
$to_go = $ARGV[0];
foreach $i (@list){
    if($i eq $to_go){
        $mark = 1;
    }
}

die "I don't know where to go, please add it or check your type!\n" if $mark == 0;
$suffix1 = "~/Dropbox/";
to_go_switch($to_go);

sub check_input {

}

# one controler to control other call to other functions
# This controler could also open the director directly
sub to_go_switch{
    my $to_go_to = $_[0];
    my $file;
    my $to_add;
    if(($to_go_to eq "study") || ($to_go_to eq "note") || ($to_go_to eq "ielts") || ($to_go_to eq "code")) {
        $file = "$suffix1$to_go_to";
    } elsif ($to_go_to eq "movie"){
        $file = "/Applications/XBMC.app/";
    }
#    `open file` if(! -d $file);
    `open $file`;
    # $file =~ s /~/\/Users\/Cao/;
    # print $file;
    # if(-d "$file") {
    #     print "yes";
    #     chdir($file);
    #     @allFiles =  `find .`;
    #     print @allFiles;
    #}
# print "opening the file or directory: $file\n";
}
