push:
	git add .
	git commit -m "done"
	git push
create:
	mkdir $(NAME)
	nano $(NAME)/solution.go
