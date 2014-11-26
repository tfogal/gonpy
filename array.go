/* Wrappers for numpy routines. */
package gonpy

/*
#define NPY_NO_DEPRECATED_API NPY_1_7_API_VERSION
#define NO_IMPORT_ARRAY
#define PY_ARRAY_UNIQUE_SYMBOL gonpy_ARRAY_API
#include <Python.h>
#include <numpy/arrayobject.h>

// This gets around PyArray_SimpleNewFromData actually being a macro, and thus
// not visible as a C.[...]-callable thing.
static PyObject*
npy_SimpleNewFromData(int nd, npy_intp* dims, int typenum, void* data) {
  return PyArray_SimpleNewFromData(nd, dims, typenum, data);
}
*/
import "C"
import(
  "errors"
  "unsafe"
  "github.com/sbinet/go-python"
)

func Array_SimpleNewFromData(ndims uint, dims []int, typenum NPY_TYPE,
                             ptr unsafe.Pointer) (*python.PyObject, error) {
  arr := C.npy_SimpleNewFromData(C.int(ndims), 
                                 ((*C.npy_intp)(unsafe.Pointer(&dims[0]))),
                                 C.int(typenum), ptr)
  if arr == nil {
    return nil, errors.New("array creation failed")
  }
  return python.PyObject_FromVoidPtr(unsafe.Pointer(arr)), nil
  //return python.Togo(arr), nil
}
