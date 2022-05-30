# Flip A Coin
## Simple command-line tool made using GO
### Monologue
I have a habit of flipping coin when I am unable to make decisions. Most of the time the heads represent the positive feeling or doing something and tails represent the negative feeling or not doing something. But, flipping the coin once and making decision based on that doesn't feel good. So my decision is always based on the frequency of heads and tails. Flipping a coin multiple time in google search is time consuming and most of the time I am in terminal. So I felt the need to make my own Flip A Coin command-line tool.

### Usage
Clone the github repo in your local machine.
```
$ https://github.com/cruzelx/Flip-A-Coin-GOLANG
```
CD into the project directory.
```
$ cd <project-directory>
```
You can run the program in two different ways. For a single toss, run the program without any flags. For repeated toss, use -R flag with integer value as the repeat value

**For single toss**
```
$ go run main.go
```

**For multiple toss**
```
$ go run main.go -R 100
```
> The coin is flipped 100 times and the aggregated result is printed in console.

### If we are the same
Build a binary using `build` command.
```
$ go build
```
Make a bin directory at HOME
```
$ mkdir ${HOME}/bin
```

Copy `flipacoin` binary to `${HOME}/bin`
```
$ cp <path-to-flipacoin-binary> ${HOME}/bin
```

Add following in shell config file. Eg, `.bashrc` or `.zshrc`

`export PATH=${HOME}/bin`

Refresh the terminal either by closing and reopening, or by using source command.
```
$ source ~/.zshrc
``` 
or
```
$ source ~/.bashrc
```

**Flip a coin**
```
$ flipacoin
```
Output: `You have got a HEAD%`
or 
```
$ flipacoin -R 40
```
Output: `Total toss: 40, Heads: (18, 45%), Tails: (22, 55%)%`


