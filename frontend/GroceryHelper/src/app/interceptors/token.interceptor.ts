import { Injectable } from '@angular/core';
import {
 HttpRequest,
 HttpHandler,
 HttpEvent,
 HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../services/auth.service';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {

 constructor(
  private authService: AuthService
 ) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    const authUser = JSON.parse(localStorage.getItem('authUser') || '{}');
    if (authUser.token) {
      request = request.clone({
        setHeaders: {
          Authorization: `Token ${authUser.token}`
        }
      });
    }
    return next.handle(request);
  }
}