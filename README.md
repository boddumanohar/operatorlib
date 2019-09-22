# Operator Library

While working on many Kubernetes operators, I realised that a lot of
code is repetative across operators. Generating, creating, updating,
deleting objects is a common thing and yet the code is being repeated
in every operator. I decided to work on project which removes the the
repeatative code from all Kubernetes Operators. The benefits of this
approach are that common well tested and stable functions can be used
by all operators. Also, this kind of tries to reduce the complexity of
dealing with Kubernetes objects which (hopefully!) will lead to more
and more vendors building there own operators. This project also
attempts to reduce the overall work required to build the operator
which means operators for small projects can be build quickly.