import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HttpHeaders } from '@angular/common/http';
import { HttpClient, HttpRequest } from '@angular/common/http';

@Component({
  selector: 'app-event-create',
  templateUrl: './event-create.component.html',
  styleUrls: ['./event-create.component.css']
})
export class EventCreateComponent implements OnInit {
  registerForm: FormGroup;
    submitted = false;
    apiURL = "http://34.70.59.156/api/createevent?text=";

    constructor(private formBuilder: FormBuilder, private httpClient: HttpClient) { }

    ngOnInit() {
        this.registerForm = this.formBuilder.group({
            title: ['', Validators.required],
            firstName: ['', Validators.required],
            lastName: ['', Validators.required],
            email: ['', [Validators.required, Validators.email]],
            location: ['', [Validators.required]],
            start: ['', [Validators.required]],
            end: ['', [Validators.required]],
        });
    }

    // convenience getter for easy access to form fields
    get f() { return this.registerForm.controls; }

    onSubmit() {
        this.submitted = true;

        // stop here if form is invalid
        if (this.registerForm.invalid) {
            return;
        }

        // display form values on success

        console.log(this.apiURL + this.registerForm.value.title + "+2+3");

        return this.httpClient.post(this.apiURL + this.registerForm.value.title + "+2+3", {}).subscribe(
            (response) => {
                console.log(response);
            }
        );
    }

    onReset() {
        this.submitted = false;
        this.registerForm.reset();
    }
}

