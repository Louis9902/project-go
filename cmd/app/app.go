package main

import (
	"modlib/internal/app"
	"modlib/internal/pkg"
	"modlib/pkg"
)

func main() {
	// app main
	modlib.Execute()
	require.RequiresApi()
	public.PublicApi()
}
