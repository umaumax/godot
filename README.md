# godot

## how to install
```
go get -u github.com/umaumax/godot
```

### before
```
digraph "Call graph" {
	label="Call graph";

	Node0x7ff45ce1a9f0 [shape=record,label="{external node}"];
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1aa80;
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1ab00;
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1abc0;
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1ac20;
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1ad40;
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1ada0;
	Node0x7ff45ce1a9f0 -> Node0x7ff45ce1aef0;
	Node0x7ff45ce1aa80 [shape=record,label="{_Z1Av}"];
	Node0x7ff45ce1aa80 -> Node0x7ff45ce1ab00;
	Node0x7ff45ce1ab00 [shape=record,label="{_Z1Bv}"];
	Node0x7ff45ce1ab00 -> Node0x7ff45ce1abc0;
	Node0x7ff45ce1ab00 -> Node0x7ff45ce1ac20;
	Node0x7ff45ce1abc0 [shape=record,label="{_Z1Cv}"];
	Node0x7ff45ce1abc0 -> Node0x7ff45ce1ad40;
	Node0x7ff45ce1ac20 [shape=record,label="{_Z1Dv}"];
	Node0x7ff45ce1ac20 -> Node0x7ff45ce1ada0;
	Node0x7ff45ce1ad40 [shape=record,label="{_Z1Ev}"];
	Node0x7ff45ce1ada0 [shape=record,label="{_Z1Fv}"];
	Node0x7ff45ce1aef0 [shape=record,label="{main}"];
	Node0x7ff45ce1aef0 -> Node0x7ff45ce1aa80;
}
```

### after
```
digraph "Call graph" {
	label="Call graph";

	"external node" [shape=record,label="{external node}"];
	"external node" -> "_Z1Av";
	"external node" -> "_Z1Bv";
	"external node" -> "_Z1Cv";
	"external node" -> "_Z1Dv";
	"external node" -> "_Z1Ev";
	"external node" -> "_Z1Fv";
	"external node" -> "main";
	"_Z1Av" [shape=record,label="{_Z1Av}"];
	"_Z1Av" -> "_Z1Bv";
	"_Z1Bv" [shape=record,label="{_Z1Bv}"];
	"_Z1Bv" -> "_Z1Cv";
	"_Z1Bv" -> "_Z1Dv";
	"_Z1Cv" [shape=record,label="{_Z1Cv}"];
	"_Z1Cv" -> "_Z1Ev";
	"_Z1Dv" [shape=record,label="{_Z1Dv}"];
	"_Z1Dv" -> "_Z1Fv";
	"_Z1Ev" [shape=record,label="{_Z1Ev}"];
	"_Z1Fv" [shape=record,label="{_Z1Fv}"];
	"main" [shape=record,label="{main}"];
	"main" -> "_Z1Av";
}
```

## TODO
* add `-i` replace option

## NOTE
* maybe `<` and `>` must be escaped

