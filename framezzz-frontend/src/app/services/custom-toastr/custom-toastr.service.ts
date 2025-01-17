import { inject, Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';

@Injectable({
  providedIn: 'root',
})
export class CustomToastrService {
  private readonly toastrService = inject(ToastrService);

  private readonly toastSettings = {
    timeOut: 3000,
    toastClass: 'toast-new',
    positionClass: 'toast-bottom-right',
    progressBar: true,
  };

  public success(message: string, title: string = 'Success'): void {
    this.toastrService.success(message, title, this.toastSettings);
  }

  public error(message: string, title: string = 'Error'): void {
    this.toastrService.error(message, title, this.toastSettings);
  }

  public warning(message: string, title: string = 'Warning'): void {
    this.toastrService.warning(message, title, this.toastSettings);
  }

  public info(message: string, title: string = 'Info'): void {
    this.toastrService.info(message, title, this.toastSettings);
  }
}
