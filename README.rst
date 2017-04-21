========
TeleSign
========

TeleSign provides the world’s most comprehensive approach to account security for Web and mobile applications.

For more information about TeleSign, visit the `TeleSign website <http://www.TeleSign.com>`_.

TeleSign REST API: Golang SDK
-----------------------------

**TeleSign web services** conform to the `REST Web Service Design Model
<http://en.wikipedia.org/wiki/Representational_state_transfer>`_. Services are exposed as URI-addressable resources
through the set of *RESTful* procedures in our **TeleSign REST API**.

The **TeleSign Golang SDK** is a set modules and functions — a *Golang Library* that wraps the
TeleSign REST API, and it simplifies TeleSign application development in the `Golang programming language
<https://www.golang.org>`_. The SDK software is distributed on
`GitHub <https://github.com/TeleSign/golang_telesign>`_.

Documentation
-------------

Detailed documentation for TeleSign REST APIs is available in the `Developer Portal <https://developer.telesign.com/>`_.

Installation
------------

To install the TeleSign Golang SDK:

.. code-block:: bash

    $ go get github.com/telesign/golang_telesign


Golang Code Example: Messaging
------------------------------

Here's a basic code example with JSON response.

.. code-block:: Golang

    package main

    import (
        ts "github.com/telesign/rest"
    )

    func main() {
        customer_id := "customer_id"
        secret_key := "secret_key"

        phone_number := "phone_number"
        message := "You're scheduled for a dentist appointment at 2:30PM."
        message_type := "ARN"
        var params map[string]string

        ts.SetCustomerID(customer_id)
        ts.SetSecretKey(secret_key)
        response = ts.Message(phone_number, message, message_type, params)
    }

.. code-block:: javascript
    
    {'errors': [],
     'reference_id': 'DGFDF6E11AB86303ASDFD425BE00000657',
     'resource_uri': '/v1/verify/DGFDF6E11AB86303ASDFD425BE00000657',
     'status': {'code': 203,
        'description': 'Delivered to gateway',
        'updated_on': '2017-02-12T00:39:58.325559Z'},
     'sub_resource': 'message',
     'verify': {'code_state': 'UNKNOWN', 'code_entered': ''}}

For more examples, see the examples folder or visit `TeleSign Developer Portal <https://developer.telesign.com/>`_.

Authentication
--------------

You will need a Customer ID and API Key in order to use TeleSign’s REST API. If you are already a customer and need an
API Key, you can generate one in the  `Portal <https://portal.telesign.com>`_.

Testing
-------

The easiest way to run the tests is to run (**go test**). Tests are located in the *src/test/* directory.

Examples
--------

Basic examples of usage are available in the directory *src/examples/*. Although the actual SDK does nt require you to 
install a dependency, examples rely on jsonparser. Available at `JSON Parser <https://github.com/buger/jsonparser>`_

.. code-block:: bash

    $ go get github.com/buger/jsonparser