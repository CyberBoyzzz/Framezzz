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
import { Router } from '@angular/router';
import { UserData } from '../../interfaces/user-data.interface';
import { CustomToastrService } from '../../services/custom-toastr/custom-toastr.service';

@Component({
  selector: 'app-register',
  imports: [
    ToolbarComponent,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.sass',
})
export class RegisterComponent {
  private readonly router = inject(Router);

  private authService = inject(AuthService);

  private customToastrService = inject(CustomToastrService);

  protected registerForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),

    password: new FormControl('', [
      Validators.required,
      Validators.minLength(8),
    ]),

    repeatPassword: new FormControl('', [
      Validators.required,
      Validators.minLength(8),
    ]),
  });

  protected submit(): void {
    if (this.registerForm.invalid) return;

    const email = this.registerForm.get('email')?.value;

    const password = this.registerForm.get('password')?.value;

    const repeatPassword = this.registerForm.get('repeatPassword')?.value;

    if (password !== repeatPassword) {
      this.customToastrService.error('Passwords do not match');

      return;
    }

    const user: UserData = { email: email!, password: password! };

    this.authService.registerUser(user).subscribe({
      next: () => {
        this.registerForm.reset();

        this.customToastrService.success(
          'Registration successful! You can now log in.'
        );

        this.router.navigate(['login']);
      },
      error: () => {
        this.customToastrService.error('Failed to register user.');
      },
    });
  }
}
