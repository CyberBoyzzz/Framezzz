import { Component, inject } from '@angular/core';
import { ToolbarComponent } from '../../components/toolbar/toolbar.component';
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AuthService } from '../../services/auth/auth.service';
import { CookieService } from 'ngx-cookie-service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  imports: [
    ToolbarComponent,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.sass',
})
export class LoginComponent {
  private readonly router = inject(Router);

  private readonly cookieService = inject(CookieService);

  private readonly authService = inject(AuthService);

  protected loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required]),
  });

  protected submit() {
    if (this.loginForm.invalid) return;

    const email = this.loginForm.get('email')?.value;
    const password = this.loginForm.get('password')?.value;

    this.authService.loginUser(email!, password!).subscribe(
      (response) => {
        this.cookieService.set('authToken', response.token, 7);

        this.router.navigate(['']);
      },
      () => {
        alert('Failed to login user.');
      }
    );
  }
}
