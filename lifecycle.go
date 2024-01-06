package httx

import "log"

func (d *Directory) Start() error {
	log.Println("env", d.secrets)
	h := d.Open("_lifecycle/start.sh")
	if h == nil {
		return nil
	}

	log.Println("=== Starting ===")
	res, err := h.Exec(nil, nil)
	if err != nil {
		return err
	}

	log.Println(res.Body)
	log.Println("=== Finished ===")
	return nil
}
