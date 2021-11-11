package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.websocket.server.PathParam;
import java.util.UUID;

@RestController
public class DemoController {
    @Autowired
    private CustomerRepository customerRepository;

    @GetMapping("/hello")
    public String sayHello(@RequestParam(value = "myName", defaultValue = "HiHiHeHe") String name) {
        String uuid = UUID.randomUUID().toString().replace("-", "").toString();
        System.out.println("YYID \t =>" + uuid);
        return String.format("Hello %s!", name + "\t your ID -->\t" + uuid);
    }

    @GetMapping("/list")
    public Iterable<Customer> getCustomers() {
        return customerRepository.findAll();
    }

    @PostMapping("/add")
    public String addCustomer(@RequestParam String first, @RequestParam String last) {
 Customer customer = new Customer();
 customer.setFirstName(first);
 customer.setLastName(last);
        customerRepository.save(customer);
        return "Add Customer success!";
    }

    @GetMapping("/find/{id}")
    public Customer findCustomerByID(@PathVariable Integer id) {
        return customerRepository.findCustomerById(id);
    }
}
