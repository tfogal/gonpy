package gonpy

/*
#define NPY_NO_DEPRECATED_API NPY_1_7_API_VERSION
#define PY_ARRAY_UNIQUE_SYMBOL gonpy_ARRAY_API
#include <Python.h>
#include <numpy/arrayobject.h>

// export
static void
gonpy_import_array() {
  _import_array();
}
*/
import "C"

// Import makes the NumPy 'import_array' function available to extension
// authors.  You need to call this function exactly once, before using any
// other NumPy functions.  Do note that Python must be initialized before this
// will work.
func Import() {
  C.gonpy_import_array()
}
