# ArgumentsAsStruct

A simple library to get the program arguments via a struct.\
After using the flag package for a long time, I started to get tired of it.\
With many possible arguments, I had to declare my variables, then apply a flag on each of them, and it got annoying.

With this library, I just have to declare a struct with my variables, add some annotations, and that's all.
# Example
```
package main

import (
    "github.com/k0kubun/pp/v3"
    Parser "github.com/mmaFR/ArgumentAsStruct"
)

type MyArguments struct {
    StrArg1 string              `default:"my default value" usage:"the argument purpose" alias:"sa1"`
    StrAgr2 string              `default:"my default value" usage:"the argument purpose" alias:"sa2"`
    BooleanArg bool             `default:"false" usage:"the argument purpose" alias:"ba" `
    IntegerArg1 int             `default:"100" usage:"the argument purpose" alias:"ia"`
    UnsignedInteger8Arg uint8   `default:"5" usage:"verbosity (0 ... 8)" alias:"v"`
}

func main() { 
    var args MyArguments
    Parser.Parse(&args)
    pp.Println(args)
}
```
Then try to start your program with `myExample -StrArg1 Hello -sa2 world -ba -IntegerArg2 9564 -v 2`

You should get this output :
```
main.MyArguments{
  StrArg1:             "Hello",
  StrAgr2:             "world",
  BooleanArg:          true,
  IntegerArg1:         100,
  IntegerArg2:         9564,
  UnsignedInteger8Arg: 2,
}
```