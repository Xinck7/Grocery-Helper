import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { map, Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';
import { LoggedInUser } from './auth'

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  baseUrl = environment.base_url
 
  login(data: any): Observable<LoggedInUser> {
    return this.http.post(`${this.baseUrl}/api-user-login/`, data)
      .pipe(map(result => {
        localStorage.setItem('authUser', JSON.stringify(result));
        return result;
      })) as Observable<LoggedInUser>; ;
  }     

  register(data: any) {
    return this.http.post(`${this.baseUrl}/api/v1/users/`, data);
  }

  getProfile(userId: string|null): Observable<any> {
    return this.http.get(`${this.baseUrl}/api/v1/users/${userId}/`);
  }

  logout() {
    localStorage.removeItem('authUser');
  }

  getAuthUser() {
    return JSON.parse(localStorage.getItem('authUser') as string);
  }
  
  setLoggedInUser(authUser: LoggedInUser): void {
    if (localStorage.getItem('authUser') !== JSON.stringify(authUser)) {
      localStorage.setItem('authUser', JSON.stringify(authUser));
    }
   }
}
