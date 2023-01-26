import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { AuthService } from "../../services/auth.service";
import { ActivatedRoute, Router } from "@angular/router";
import { ToastrService } from 'ngx-toastr';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup;
  returnUrl!: string;
  isSubmitted: boolean = false;

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private toastr: ToastrService,
    private route: ActivatedRoute,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'];
  }
  
  onSubmit() {
    this.isSubmitted = true;
    this.authService.login(this.loginForm.value).subscribe({
      next: (data: any) => {
        console.log("data before setLoggedInUser data id", data.id);
        this.authService.setLoggedInUser(data);
        console.log("data in login auth service in the submit before next", data);
        this.router.navigateByUrl(`/user-profile/${data.id}`);
      },
      error: (error) => {
        console.log(error);
        this.toastr.error(error);
      }
    })
    .add(() => {
      this.isSubmitted = false;
    });
  }
}