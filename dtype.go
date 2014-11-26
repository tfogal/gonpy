package gonpy

/*
#define NPY_NO_DEPRECATED_API NPY_1_7_API_VERSION
#define PY_ARRAY_UNIQUE_SYMBOL gonpy_ARRAY_API
#define NO_IMPORT_ARRAY
#include <Python.h>
#include <numpy/arrayobject.h>
*/
import "C"

type NPY_TYPE C.int
const(
  NPY_BOOL NPY_TYPE = C.NPY_BOOL
  NPY_BYTE = iota
  NPY_UBYTE
  NPY_SHORT
  NPY_USHORT
  NPY_INT
  NPY_UINT
  NPY_LONG
  NPY_ULONG
  NPY_LONGLONG
  NPY_ULONGLONG
  NPY_FLOAT
  NPY_DOUBLE
  NPY_LONGDOUBLE
  NPY_CFLOAT
  NPY_CDOUBLE
  NPY_CLONGDOUBLE
  NPY_OBJECT=17
  NPY_STRING
  NPY_UNICODE
  NPY_VOID
  /*
   * New 1.6 types appended, may be integrated
   * into the above in 2.0.
   */
  NPY_DATETIME
  NPY_TIMEDELTA
  NPY_HALF

  NPY_NTYPES
  NPY_NOTYPE
  NPY_CHAR      /* special flag */
  NPY_USERDEF=256  /* leave room for characters */

  /* The number of types not including the new 1.6 types */
  NPY_NTYPES_ABI_COMPATIBLE=21
)

// These are defines massaged from 'ndarraytypes.h'.
const(
  NPY_ARRAY_C_CONTIGUOUS = 0x0001
  NPY_ARRAY_ALIGNED = 0x100
  NPY_ARRAY_WRITEABLE = 0x400
  NPY_ARRAY_BEHAVED = NPY_ARRAY_ALIGNED | NPY_ARRAY_WRITEABLE
  NPY_ARRAY_CARRAY = NPY_ARRAY_C_CONTIGUOUS | NPY_ARRAY_BEHAVED
)
