package multimedia

import (
	"strconv"
	"strings"
)

type ContenidoMultimedia struct {
	Multimedias []Multimedia
}

func (cm *ContenidoMultimedia) Mostrar() string {
	var datos string
	for _, m := range cm.Multimedias {
		datos += strings.Repeat("-", 65) + "\n"
		datos += m.Mostrar() + "\n"
	}
	datos += strings.Repeat("-", 65) + "\n"
	return datos
}

func (cm *ContenidoMultimedia) AgregarMultimedia(m Multimedia) {
	cm.Multimedias = append(cm.Multimedias, m)
}

type Multimedia interface {
	Mostrar() string
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales int64
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int64
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (v *Video) Mostrar() string {
	return "| Titulo: " + v.Titulo + "\t| Formato: " + v.Formato + "\t| Frames: " + strconv.FormatInt(v.Frames, 10) + " frames\t|" //Base 10
}

func (a *Audio) Mostrar() string {
	return "| Titulo: " + a.Titulo + "\t| Formato: " + a.Formato + "\t| Duracion: " + strconv.FormatInt(a.Duracion, 10) + " seg\t|"
}

func (i *Imagen) Mostrar() string {
	return "| Titulo: " + i.Titulo + "\t| Formato: " + i.Formato + "\t| Canales: " + strconv.FormatInt(i.Canales, 10) + " canales\t|"
}
