package diccionario

import (
	"fmt"
	Hash "hash/fnv"
	TDALista "tdas/lista"
)

const (
	_CAPACIDAD_INICIAL = 5
)

type parClaveValor[K comparable, V any] struct {
   clave K
   dato  V
}

type diccionarioHash[K comparable, V any] struct {
   tabla    []TDALista.Lista[parClaveValor[K,V]]
   tam      int
   cantidad int
}

type iteradorHash[K comparable, V any] struct {
	dic       *diccionarioHash[K, V]
	posTabla  int
	iterLista TDALista.iteradorLista[parClaveValor[K, V]]
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (d *diccionarioHash[K,V]) obtenerIndice (clave T){
	d := fnv.New32a()
	d.Write(convertirABytes(clave))
	return int(d.Sum32() % uint32(len(d.tabla)))
}

func crearTablaDic[K comparable, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	dic := new(diccionarioHash[K,V])
	dic.tabla = crearTabla[K, V](_CAPACIDAD_INICIAL)
	dic.tam = _CAPACIDAD_INICIAL
	return dic
}

func (iter *iteradorLista[parClaveValor[K, V]]) HayAlgoMas() bool {
	return iter.iterLista.HayAlgoMas()
}

func (iter *iteradorLista[parClaveValor[K, V]]) VerActual() (K, V) {
	if !iter.iterLista.HayAlgoMas() {
		panic("a")
	}
	par := iter.iterLista.VerActual()
	return  par.clave, par.dato
}

func (iter *iteradorLista[parClaveValor[K, V]]) Siguiente() {
	
}

func (d *diccionarioHash[K,V]) Cantidad() int {
	return d.cantidad
}

//buscar clave
func buscarClave (clave K) bool {

}

func (d *diccionarioHash[K,V]) Pertence(clave K) bool {

}

func (d *diccionarioHash[K,V]) Guardar(clave K, dato V){
	if 
}

