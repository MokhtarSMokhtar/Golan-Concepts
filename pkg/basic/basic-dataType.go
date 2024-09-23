package basic

/*
1. Basic Data Types: int, float64, string, bool

// 1. Numeric Types
/*
1. Type `int`
   - **Description:** Represents signed integers.
   - **Variants:**
     - `int8`: 8-bit signed integer.
       - **Range:** -128 to 127
     - `int16`: 16-bit signed integer.
       - **Range:** -32,768 to 32,767
     - `int32`: 32-bit signed integer.
       - **Range:** -2,147,483,648 to 2,147,483,647
     - `int64`: 64-bit signed integer.
       - **Range:** -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
   - **Notes:**
     - `int` is a signed integer type that is at least 32 bits in size.
     - It is a distinct type and not an alias for any specific integer size like `int32`.

2. Type `uint`
   - **Description:** Represents unsigned integers.
   - **Variants:**
     - `uint8`: 8-bit unsigned integer.
       - **Range:** 0 to 255
     - `uint16`: 16-bit unsigned integer.
       - **Range:** 0 to 65,535
     - `uint32`: 32-bit unsigned integer.
       - **Range:** 0 to 4,294,967,295
     - `uint64`: 64-bit unsigned integer.
       - **Range:** 0 to 18,446,744,073,709,551,615
   - **Notes:**
     - `uint` is an unsigned integer type that is at least 32 bits in size.

3. Type `float64`
   - **Description:** Represents 64-bit floating-point numbers.
   - **Usage:** Used for decimal numbers requiring precision.
   - **Example Range:** Approximately ±1.8e308

4. Type `string`
   - **Description:** Represents a sequence of Unicode characters.
   - **Usage:** Used for storing and manipulating text data.

5. Type `bool`
   - **Description:** Represents boolean values.
   - **Values:** `true` or `false`
   - **Usage:** Used for conditional operations and flags.
*/
// more data types https://github.com/golang/go/blob/master/src/builtin/builtin.go
