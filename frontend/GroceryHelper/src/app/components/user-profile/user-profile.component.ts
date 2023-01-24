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
    this.fetchProfile();
  }

  fetchProfile() {
    const userId = this.activatedRoute.snapshot.paramMap.get('id');
    this.authService.profile(userId).subscribe({
      next: (result) => {
        this.user = result;
      },
      error: (error) => {
        this.toastr.error(error);
      }
    });
  }
  // userProfile: UserProfile|null = null;

  // constructor(private userProfileService: UserProfileService, private activatedRoute: ActivatedRoute) { }

  // ngOnInit(): void {
  //   const userId = this.activatedRoute.snapshot.paramMap.get('id');
  //   this.userProfileService.getUserProfile(userId).subscribe({
  //       next: (data) => {
  //         console.log(data);
  //         this.userProfile = data;
  //       },
  //       error: (error) => {
  //         console.log(error);
  //       }
  //     }
  //   );
  // }
}