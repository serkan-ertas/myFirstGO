package intface

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

func SaveData(data Saver) error {
	err := data.Save()

	if err != nil {
		return err
	}
	return nil
}

func OutputData(data Outputtable) error {
	data.Display()
	return SaveData(data)
}
