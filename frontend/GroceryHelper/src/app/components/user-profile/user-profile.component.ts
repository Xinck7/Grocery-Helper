import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { ToastrService } from 'ngx-toastr';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {

  user!: any;

  constructor(
    private authService: AuthService,
    private toastr: ToastrService,
    private activatedRoute: ActivatedRoute
  ) { }

  ngOnInit(): void {
    // this.fetchProfile();
    const userId = this.activatedRoute.snapshot.paramMap.get('id');
    this.authService.getProfile(userId).subscribe({
      next: (data) => {
        this.user = data;
      },
      error: (error) => {
        this.toastr.error(error);
      }
    });
  }
}