package diccionario

import (
	"fmt"
	Hash "hash/fnv"
	TDALista "tdas/lista"
)

const (
	_CAPACIDAD_INICIAL = 5
	_PANIC_ITERADOR = "El iterador termino de iterar"
	_PANIC_CLAVE_NO_PERTENECE = "La clave no pertenece al diccionario"
	_FACTOR_CARGA_MAX = 2.0
)

type iteradorHash[K comparable, V any] struct {
	dic       *diccionarioHash[K, V]
	posTabla  int
	iterLista TDALista.IteradorLista[parClaveValor[K, V]]
}

type parClaveValor[K comparable, V any] struct {
   clave K
   dato  V
}

type diccionarioHash[K comparable, V any] struct {
   tabla    []TDALista.Lista[parClaveValor[K,V]]
   tam      int
   cantidad int
   iterador *iteradorHash[K,V]
}



func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (d *diccionarioHash[K,V]) obtenerIndice (clave K) int {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return int(h.Sum32() % uint32(len(d.tabla)))
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
	dic.tabla = crearTablaDic[K, V](_CAPACIDAD_INICIAL)
	dic.tam = _CAPACIDAD_INICIAL
	return dic
}



// ITERADOR EXTERNO --------------------

func (iter *iteradorHash[K, V]) HayAlgoMas() bool {
	return iter.iterLista != nil && iter.iterLista.HayAlgoMas()
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_PANIC_ITERADOR)
	}
	par := iter.iterLista.VerActual()
	return  par.clave, par.dato
}

func (iter *iteradorHash[K, V]) Avanzar()  {
	if !iter.HayAlgoMas() {
		panic(_PANIC_ITERADOR)
	}
	iter.iterLista.Avanzar()
	if !iter.HayAlgoMas() {
		iter.posTabla++
		for iter.posTabla < iter.dic.tam && iter.dic.tabla[iter.posTabla].EstaVacia() {
			iter.posTabla++
		}
		if iter.posTabla < iter.dic.tam {
			iter.iterLista = iter.dic.tabla[iter.posTabla].Iterador()
		}
	}
}




// HASH -----------------------------

func (d *diccionarioHash[K, V]) buscarClave(clave K) (TDALista.IteradorLista[parClaveValor[K, V]], bool) {
	iter := d.tabla[d.obtenerIndice(clave)].Iterador()
	for iter.HayAlgoMas() {
		if iter.VerActual().clave == clave {
			return iter, true
		}
		iter.Avanzar()
	}
	return iter, false
}

func (d *diccionarioHash[K,V]) Cantidad() int {
	return d.cantidad
}

func (d *diccionarioHash[K, V]) Pertenece(clave K) bool {
	_, encontrada := d.buscarClave(clave)
	return encontrada
}

func (d *diccionarioHash[K, V]) Guardar(clave K, dato V) {
	par := parClaveValor[K, V]{clave, dato}
	if iter, encontrada := d.buscarClave(clave); encontrada {
		iter.Borrar()
		iter.Insertar(par)
		return
	}
	if float64(d.cantidad)/float64(d.tam) >= _FACTOR_CARGA_MAX {
		d.redimensionar(d.tam * 2)
	}
	d.tabla[d.obtenerIndice(clave)].InsertarUltimo(par)
	d.cantidad++
}

func (d *diccionarioHash[K,V]) Obtener(clave K) V {

func (d *diccionarioHash[K,V]) Borrar(clave K) V {

func (d *diccionarioHash[K,V]) Iterar( f func(clave K, dato V) bool) {

func (d *diccionarioHash[K,V]) Iterador() IterDiccionario[K, V] {

func (d *diccionarioHash[K,V]) redimensionar(nuevoTam int) {

func (d *diccionarioHash[K,V]) rehash() {



