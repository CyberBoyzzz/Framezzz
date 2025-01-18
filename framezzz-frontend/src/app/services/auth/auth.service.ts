import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserData } from '../../interfaces/user-data.interface';
import { LoginResponse } from '../../interfaces/login-response.interface';
import { RegisterResponse } from '../../interfaces/register-response.interface';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly httpClient = inject(HttpClient);

  public registerUser(user: UserData): Observable<RegisterResponse> {
    return this.httpClient.post<RegisterResponse>(`/auth-api/auth/register`, {
      email: user.email,
      password: user.password,
    });
  }

  public loginUser(user: UserData): Observable<LoginResponse> {
    return this.httpClient.post<LoginResponse>(`/auth-api/auth/login`, {
      email: user.email,
      password: user.password,
    });
  }
}
