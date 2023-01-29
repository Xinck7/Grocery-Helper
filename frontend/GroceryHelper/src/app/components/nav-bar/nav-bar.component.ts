import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent {
  loggedIn!: boolean;

  constructor(
    public authService: AuthService,
  ) { }


  ngOnInit(): void {
    const authUser = JSON.parse(localStorage.getItem('authUser') || '{}');
    if (authUser.token){
      this.loggedIn = true
     }
     else {
      this.loggedIn = false
     }
   }

  handleLogout() {
    this.authService.logout()
    this.loggedIn = false
  }
}
