// PlainHeader.js   functional component
import React from 'react';
import { Navbar } from 'react-bootstrap'; 
import Form from 'react-bootstrap/Form';
import 'bootstrap/dist/css/bootstrap.css';

function TestComponent() {
    return (
      <div>
        <Navbar fixed='top' expand='md' bg='info' sticky='top' >
          <Form inline> 
            <h5 style={{ color: 'black' }}><b><i>Test!</i></b></h5>
          </Form>
        </Navbar>
      </div>
    );
}

export default TestComponent;