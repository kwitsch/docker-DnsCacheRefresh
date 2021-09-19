SET DCR.resolver=docker.ddv
SET DCR.start=10s
SET DCR.refresh=1m
SET DCR.verbose=True
SET DCR.domains[0]=github.com
SET DCR.domains[1]=google.de
SET DCR.domains[2]=twitter.com

go run dnscacherefresh