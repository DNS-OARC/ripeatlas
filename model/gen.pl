#!/usr/bin/env perl

use utf8;
use common::sense;
binmode(STDOUT, ":utf8");

unless (scalar @ARGV == 1) {
    die "usage: gen.pl \$file";
}
my ($file) = @ARGV;
my $package = `pwd`;
$package =~ s/[\r\n]+$//o;
my $subpackage_path = $package;
$package =~ s/.*\///o;
$subpackage_path =~ s/.*model\/?//o;
if ($subpackage_path) {
    $subpackage_path .= "/";
}
if ($package eq "model") {
    $package = "ripeatlas";
}
my $subpackage = $file;
$subpackage =~ s/\.model$//o;
my $name = camalize($subpackage);
my $cc = substr($file, 0, 1);

unless (open(FILE, "<:encoding(UTF-8)", $file)) {
    die "open($file): $!";
}

my %param;
my @order;
my $need_subpackage;

while (<FILE>) {
    my ($section, $param, $type, $desc, $array, $obj);

    s/[\r\n]+$//o;

    if (/^\s*\[(\w+)\]\s+(\w+)\s+\(([^\)]+)\)(.*)/o) {
        ($section, $param, $type, $desc) = ($1, $2, $3, $4);
    }
    elsif (/^\s*(\w+)\s+\(([^\)]+)\)(.*)/o) {
        ($param, $type, $desc) = ($1, $2, $3);
    }
    else {
        die "unknown: $_";
    }

    $desc =~ s/^\s*=\s*\[[^\]]+\]//o;
    $desc =~ s/^[^:]*:\s*//o;
    $desc =~ s/\s*,\s*$/./o;
    unless ($desc =~ /\.$/o) {
        $desc .= ".";
    }

    if ($type =~ /array\[([^\]]+)\]/o) {
        $type = $1;
        $array = "[]";
    }

    if ($type eq "integer") {
        $type = "int";
    }
    elsif ($type eq "string") {
        $type = "string";
    }
    elsif ($type eq "boolean") {
        $type = "bool";
    }
    elsif ($type eq "float") {
        $type = "float64";
    }
    else {
        $type = "*$subpackage.".camalize($type);
        $obj = 1;
        $need_subpackage = 1;
    }
    $type = $array.$type;

    if (exists $param{$param}) {
        unless ($param{$param}->{type} eq $type) {
            die "shared param $param with incompatible type";
        }

        if (ref($param{$param}->{section}) eq 'HASH') {
            $param{$param}->{section}->{$section} = $desc;
        }
        else {
            $param{$param}->{section} = {
                $param{$param}->{section} => $param{$param}->{desc},
                $section => $desc,
            };
            delete $param{$param}->{desc};
        }
    }
    else {
        $param{$param} = {
            param => $param,
            section => $section,
            type => $type,
            desc => $desc,
            obj => $obj,
        };
        push(@order, $param);
    }
}

say "package $package";
say "";
say "import (";
say "    \"encoding/json\"";
say "    \"fmt\"";
if ($need_subpackage) {
    say "";
    say "    \"github.com/DNS-OARC/ripeatlas/$subpackage_path$subpackage\""
}
say ")";
say "";
say "type $name struct {";
say "    ParseError error";
say "";
say "    data struct {";
foreach (@order) {
    my $p = $param{$_};
    if ($p->{obj}) {
        say "        ".camalize($p->{param})." json.RawMessage `json:\"$p->{param}\"`"
    }
    else {
        say "        ".camalize($p->{param})." $p->{type} `json:\"$p->{param}\"`"
    }
}
say "    }";
foreach (@order) {
    my $p = $param{$_};
    unless ($p->{obj}) {
        next;
    }

    say "";
    say "    ".lcfirst(camalize($p->{param}))." $p->{type}"
}
say "}";
say "";
say "func ($cc *$name) UnmarshalJSON(b []byte) error {";
say "    if err := json.Unmarshal(b, &$cc.data); err != nil {";
say "        return fmt.Errorf(\"%s for %s\", err.Error(), string(b))";
say "    }";
foreach (@order) {
    my $p = $param{$_};
    unless ($p->{obj}) {
        next;
    }
    say "";
    say "    if $cc.data.".camalize($p->{param})." != nil {";
    say "        if err := json.Unmarshal($cc.data.".camalize($p->{param}).", &$cc.".lcfirst(camalize($p->{param}))."); err != nil {";
    say "            return fmt.Errorf(\"Unable to process $name ".camalize($p->{param}).": %s\", err.Error())";
    say "        }";
    say "    }";
}
say "";
say "    return nil";
say "}";
foreach (@order) {
    my $p = $param{$_};
    say "";
    if (ref($p->{section}) eq 'HASH') {
        foreach my $s (sort keys %{$p->{section}}) {
            say "// [$s] $p->{section}->{$s}";
        }
    }
    elsif ($p->{section}) {
        say "// [$p->{section}] $p->{desc}";
    }
    else {
        say "// $p->{desc}";
    }
    say "func ($cc *$name) ".camalize($p->{param})."() $p->{type} {";
    if ($p->{obj}) {
        say "    return $cc.".lcfirst(camalize($p->{param}));
    }
    else {
        say "    return $cc.data.".camalize($p->{param});
    }
    say "}";
}

close(FILE);

sub camalize {
    my ($s) = @_;

    return join("", map { ucfirst($_) } split(/[\W_]+/o, $s));
}
