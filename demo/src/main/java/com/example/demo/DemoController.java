package com.example.demo;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
public class DemoController {
    @GetMapping("/hello")
    public String generateString(@RequestParam(value = "myName", defaultValue = "HihiHehe") String name) {
        String uuid = UUID.randomUUID().toString();
        System.out.println(uuid);
        return String.format("Helloo %s!", name + "\tyour ID ->\t" + uuid);
    }
}
