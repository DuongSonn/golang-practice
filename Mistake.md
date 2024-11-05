# Golang Mistake Notes

-   Don't abuse global variable inside a package. Because function inside the package can change that variable
-   Don't use too much interface. Use cases for interface:
    -   Common behavior: many types have the same behavior
    -   Decoupling: Abstract help the implementation can be replaced without affecting the other code. Also benefit unit test (you can mock the method)
    -   Restricting behavior: prevent updating data of a `struct`
