package manifest

type Manifest struct {
	Title string
}

func Create(title string) Manifest {
    return Manifest{title}
}